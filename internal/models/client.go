package models

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/charmbracelet/log"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/hhirsch/builder/internal/helpers/system"
	"github.com/melbahja/goph"
	"golang.org/x/crypto/ssh"
	"os"
	"os/user"
	"strings"
	"time"
)

type Client struct {
	sshClient     goph.Client
	keyPath       string
	user          string
	host          string
	targetUser    string
	logger        *helpers.Logger
	step          string
	stepNumber    int
	commandNumber int
}

func NewClient(environment *Environment, userName string, host string) *Client {
	currentUser, err := user.Current()
	keyPath := currentUser.HomeDir + "/.ssh/id_rsa"

	logger := environment.GetLogger()
	client := &Client{
		user:    userName,
		host:    host,
		logger:  logger,
		keyPath: keyPath,
	}

	if err != nil {
		logger.Fatal(err)
	}

	client.ensureSnapshotDirectoryExists()

	auth, err := goph.Key(client.keyPath, "")
	if err != nil {
		logger.Fatal(err)
	}

	sshClient, err := goph.New(userName, host, auth)
	if err != nil {
		logger.Fatal(err)
	}

	client.sshClient = *sshClient
	return client
}

func (this *Client) ensureSnapshotDirectoryExists() {
	path := "snapshots"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			this.logger.Warn(err.Error())
		}
	}
}

func (this *Client) Execute(command string) string {
	out, error := this.sshClient.Run(command)

	if error != nil {
		this.logger.Warn(error.Error())
	}
	return string(out)
}

func (this *Client) ExecuteAndPrint(command string) {
	log.Info("Executing " + command)
	fmt.Println(this.Execute(command))
}

func (this *Client) EnsurePath(path string) {
	_, error := this.sshClient.Run("ls " + path)
	if error != nil {
		this.logger.Warn(error.Error())
		if len(this.targetUser) > 0 {
			this.Execute("sudo -u " + this.targetUser + " mkdir -p " + path)
		} else {
			this.Execute("mkdir -p " + path)
		}
	}
}

func (this *Client) PushFile(source string, target string) {
	this.logger.Info("Source for the upload: " + source)
	var adaptedTarget = target
	if len(this.targetUser) > 0 && !strings.HasPrefix(target, "/") {
		adaptedTarget = "/home/" + this.targetUser + "/" + target
	}
	lastIndex := strings.LastIndex(adaptedTarget, "/")
	if lastIndex == -1 {
		this.logger.Fatal("Target must contain a slash.")
		return
	}

	path := adaptedTarget[:lastIndex]
	this.EnsurePath(path)
	this.logger.Info("Uploading file from: " + source + " to: " + adaptedTarget)
	this.Upload(source, adaptedTarget)
	this.logger.Info("File Uploaded")
	if len(this.targetUser) > 0 {
		this.SetFileOwner(adaptedTarget, this.targetUser)
	}
}

func (this *Client) UploadSftp(source string, target string) {
	err := this.sshClient.Upload(source, target)
	if err != nil {
		this.logger.Fatal("Error uploading: " + err.Error())
	}
}

// Gives a binary file permission to open network ports
func (this *Client) EnsureCapabilityConnection(path string) {
	if len(this.targetUser) > 0 && !strings.HasPrefix(path, "/") {
		this.ExecuteAndPrint("setcap 'cap_net_bind_service=+ep' /home/" + this.targetUser + "/" + path)
		return
	}
	this.ExecuteAndPrint("setcap 'cap_net_bind_service=+ep' " + path)
}

func (this *Client) Upload(source string, target string) {
	clientConfig, _ := auth.PrivateKey(this.user, this.keyPath, ssh.InsecureIgnoreHostKey())
	client := scp.NewClient(this.host+":22", &clientConfig)
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}
	f, _ := os.Open(source)
	defer client.Close()
	defer f.Close()
	err = client.CopyFromFile(context.Background(), *f, target, "0775")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}

func (this *Client) EnsureService(serviceName string, path string, description string) {
	if len(this.targetUser) <= 0 {
		this.logger.Fatal("ensureService requires a target user to be set run the command setTargetUser")
	}
	this.EnsureCustomService(serviceName, this.targetUser, "/home/"+this.targetUser+"/"+path, description)
}

func (this *Client) EnsureCustomService(serviceName string, userName string, path string, description string) {
	systemd := system.Systemd{}
	config := systemd.GetConfig(userName, path, description)

	hash := md5.New()
	hash.Write([]byte(config))
	hashSum := hash.Sum(nil)
	tempFilePath := "/tmp/builder-" + hex.EncodeToString(hashSum)

	err := os.WriteFile(tempFilePath, []byte(config), 0644)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
		return
	}
	this.logger.Info("Temporary file" + tempFilePath + " for service config created.")
	this.logger.Info("Uploading file")
	this.Upload(tempFilePath, "/etc/systemd/system/"+serviceName+".service")
	this.logger.Info("Reloading systemd config")
	this.ExecuteAndPrint("systemctl daemon-reload")
	this.logger.Info("Enable service on start " + serviceName)
	this.ExecuteAndPrint("systemctl enable " + serviceName)
	this.logger.Info("Start service " + serviceName)
	this.ExecuteAndPrint("systemctl start " + serviceName)
	this.logger.Info("Service status " + serviceName)
	this.ExecuteAndPrint("systemctl status " + serviceName)
}

func (this *Client) SetTargetUser(userName string) {
	this.targetUser = userName
	this.ensureUserExists(userName)
}

func (this *Client) SetFileOwner(path string, userName string) {
	this.logger.Info("Ensure file " + path + " belongs to user " + userName + ".")
	result := this.Execute("chown -R " + userName + ":" + userName + " " + path)
	if len(result) == 0 {
		this.logger.Info("Success")
	} else {
		this.logger.Info("Setting user permission")
	}
}

func (this *Client) EnsureExecutable(path string) {
	this.logger.Info("Ensure file is Executable ")
	if len(this.targetUser) > 0 && !strings.HasPrefix(path, "/") {
		this.Execute("chmod +x " + "/home/" + this.targetUser + "/" + path)
		this.logger.Info("Path: " + "/home/" + this.targetUser + "/" + path)
		return
	}
	this.Execute("chmod +x " + path)
	this.logger.Info("Path: " + path)
}

func (this *Client) ensureUserExists(userName string) {
	this.logger.Info("Checking for user " + userName)
	result := this.Execute("getent passwd " + userName)
	if len(result) == 0 {
		this.createUser(userName)
	} else {
		this.logger.Info("User " + userName + " exists.")
	}
}

func (this *Client) createUser(userName string) {
	this.logger.Info("User " + userName + " does not exist.")
	result := this.Execute("useradd -m " + userName)
	if len(result) == 0 {
		this.logger.Info("User " + userName + " created successfully.")
	} else {
		this.logger.Fatal("Error during User creation")
	}
}

func (this *Client) EnsurePackage(packageName string) {
	this.logger.Info("Checking status of package " + packageName)
	this.logger.Info("Status of " + packageName + " is not installed")
	this.ExecuteAndPrint("dpkg --status " + packageName)
	this.logger.Info("Installing " + packageName)
	this.ExecuteAndPrint("apt-get update")
	this.ExecuteAndPrint("apt-get install " + packageName)
}

func (this *Client) DumpPackages() {
	this.logger.Info("Dumping Packages")
	currentTime := time.Now()
	fileName := "snapshots/" + currentTime.Format("02-01-2006_15-04-05") + ".dmp" // File name format: DD-MM-YYYY_HH-MM-SS

	err := os.WriteFile(fileName, []byte(this.Execute("dpkg --get-selections")), 0644)
	if err != nil {
		this.logger.Fatal(fmt.Printf("Error writing file: %v", err))
	}

	this.logger.Info("File " + fileName + " created and string written successfully!\n")
}

func (this *Client) ListPackages() {
	this.logger.Info("Listing Packages")
	this.ExecuteAndPrint("dpkg --get-selections")
}

func (this *Client) TearDown() {
	defer this.sshClient.Close()
}
