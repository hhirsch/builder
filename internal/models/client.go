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
)

type Client struct {
	sshClient  goph.Client
	keyPath    string
	user       string
	host       string
	targetUser string
	logger     *helpers.Logger
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

func (client *Client) GetHost() (hostName string) {
	return client.host
}

func (client *Client) ensureSnapshotDirectoryExists() {
	path := "snapshots"
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			client.logger.Warn(err.Error())
		}
	}
}

func (client *Client) Execute(command string) string {
	out, error := client.sshClient.Run(command)

	if error != nil {
		client.logger.Warn(error.Error())
	}
	return string(out)
}

func (client *Client) ExecuteAndPrint(command string) {
	log.Info("Executing " + command)
	fmt.Println(client.Execute(command))
}

func (client *Client) EnsurePath(path string) {
	_, error := client.sshClient.Run("ls " + path)
	if error != nil {
		client.logger.Warn(error.Error())
		if len(client.targetUser) > 0 {
			client.Execute("sudo -u " + client.targetUser + " mkdir -p " + path)
		} else {
			client.Execute("mkdir -p " + path)
		}
	}
}

func (client *Client) PushFile(source string, target string) {
	client.logger.Info("Source for the upload: " + source)
	var adaptedTarget = target
	if len(client.targetUser) > 0 && !strings.HasPrefix(target, "/") {
		adaptedTarget = "/home/" + client.targetUser + "/" + target
	}
	lastIndex := strings.LastIndex(adaptedTarget, "/")
	if lastIndex == -1 {
		client.logger.Fatal("Target must contain a slash.")
		return
	}

	path := adaptedTarget[:lastIndex]
	client.EnsurePath(path)
	client.logger.Info("Uploading file from: " + source + " to: " + adaptedTarget)
	client.Upload(source, adaptedTarget)
	client.logger.Info("File Uploaded")
	if len(client.targetUser) > 0 {
		client.SetFileOwner(adaptedTarget, client.targetUser)
	}
}

func (client *Client) UploadSftp(source string, target string) {
	err := client.sshClient.Upload(source, target)
	if err != nil {
		client.logger.Fatal("Error uploading: " + err.Error())
	}
}

// Gives a binary file permission to open network ports
func (client *Client) EnsureCapabilityConnection(path string) {
	if len(client.targetUser) > 0 && !strings.HasPrefix(path, "/") {
		client.ExecuteAndPrint("setcap 'cap_net_bind_service=+ep' /home/" + client.targetUser + "/" + path)
		return
	}
	client.ExecuteAndPrint("setcap 'cap_net_bind_service=+ep' " + path)
}

func (client *Client) Upload(source string, target string) {
	clientConfig, _ := auth.PrivateKey(client.user, client.keyPath, ssh.InsecureIgnoreHostKey())
	sshClient := scp.NewClient(client.host+":22", &clientConfig)
	err := sshClient.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}
	f, _ := os.Open(source)
	defer sshClient.Close()
	defer func() {
		err := f.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()
	err = sshClient.CopyFromFile(context.Background(), *f, target, "0775")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}

func (client *Client) EnsureService(serviceName string, path string, description string) {
	if len(client.targetUser) <= 0 {
		client.logger.Fatal("ensureService requires a target user to be set run the command setTargetUser")
	}
	client.EnsureCustomService(serviceName, client.targetUser, "/home/"+client.targetUser+"/"+path, description)
}

func (client *Client) EnsureCustomService(serviceName string, userName string, path string, description string) {
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
	client.logger.Info("Temporary file" + tempFilePath + " for service config created.")
	client.logger.Info("Uploading file")
	client.Upload(tempFilePath, "/etc/systemd/system/"+serviceName+".service")
	client.logger.Info("Reloading systemd config")
	client.ExecuteAndPrint("systemctl daemon-reload")
	client.logger.Info("Enable service on start " + serviceName)
	client.ExecuteAndPrint("systemctl enable " + serviceName)
	client.logger.Info("Start service " + serviceName)
	client.ExecuteAndPrint("systemctl start " + serviceName)
	client.logger.Info("Service status " + serviceName)
	client.ExecuteAndPrint("systemctl status " + serviceName)
}

func (client *Client) SetTargetUser(userName string) {
	client.targetUser = userName
	client.ensureUserExists(userName)
}

func (client *Client) SetFileOwner(path string, userName string) {
	client.logger.Info("Ensure file " + path + " belongs to user " + userName + ".")
	result := client.Execute("chown -R " + userName + ":" + userName + " " + path)
	if len(result) == 0 {
		client.logger.Info("Success")
	} else {
		client.logger.Info("Setting user permission")
	}
}

func (client *Client) EnsureExecutable(path string) {
	client.logger.Info("Ensure file is Executable ")
	if len(client.targetUser) > 0 && !strings.HasPrefix(path, "/") {
		client.Execute("chmod +x " + "/home/" + client.targetUser + "/" + path)
		client.logger.Info("Path: " + "/home/" + client.targetUser + "/" + path)
		return
	}
	client.Execute("chmod +x " + path)
	client.logger.Info("Path: " + path)
}

func (client *Client) ensureUserExists(userName string) {
	client.logger.Info("Checking for user " + userName)
	result := client.Execute("getent passwd " + userName)
	if len(result) == 0 {
		client.createUser(userName)
	} else {
		client.logger.Info("User " + userName + " exists.")
	}
}

func (client *Client) createUser(userName string) {
	client.logger.Info("User " + userName + " does not exist.")
	result := client.Execute("useradd -m " + userName)
	if len(result) == 0 {
		client.logger.Info("User " + userName + " created successfully.")
	} else {
		client.logger.Fatal("Error during User creation")
	}
}

func (client *Client) EnsurePackage(packageName string) {
	client.logger.Info("Checking status of package " + packageName)
	client.logger.Info("Status of " + packageName + " is not installed")
	client.ExecuteAndPrint("dpkg --status " + packageName)
	client.logger.Info("Installing " + packageName)
	client.ExecuteAndPrint("apt-get update")
	client.ExecuteAndPrint("apt-get install " + packageName)
}

func (client *Client) ListPackages() {
	client.logger.Info("Listing Packages")
	client.ExecuteAndPrint("dpkg --get-selections")
}

func (client *Client) TearDown() {
	defer func() {
		err := client.sshClient.Close()
		if err != nil {
			fmt.Printf("error closing ssh client %v", err)
		}
	}()
}
