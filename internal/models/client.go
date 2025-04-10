// nolint
// file will be sunset, soon
package models

import (
	"context"
	"errors"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"github.com/charmbracelet/log"
	"github.com/hhirsch/builder/internal/helpers"
	"github.com/melbahja/goph" //https://pkg.go.dev/github.com/melbahja/goph
	"golang.org/x/crypto/ssh"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type Client struct {
	sshClient  goph.Client
	keyPath    string
	userName   string
	host       string
	TargetUser string
	logger     *helpers.Logger
}

func NewClient(environment *Environment, userName string, host string) *Client {
	currentUser, err := user.Current()
	keyPath := currentUser.HomeDir + "/.ssh/id_rsa"
	logger := environment.GetLogger()
	client := &Client{
		userName: userName,
		host:     host,
		logger:   logger,
		keyPath:  keyPath,
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

func (client *Client) IsConnected() bool {
	//	return client.sshClient != nil
	return true
}

func (client *Client) ExecuteOnLocalhost(command string) (string, error) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "", errors.New("no command parameter set")
	}
	client.logger.Infof("Running %s on localhost.", parts[0])
	cmd := exec.Command(parts[0], parts[1:]...)

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(string(output))

	return result, nil
}

func (client *Client) Execute(command string) (string, error) {
	out, error := client.sshClient.Run(command)

	if error != nil {
		return "", error
	}
	return string(out), nil
}

func (client *Client) ExecuteAndPrint(command string) {
	log.Info("Executing " + command)
	fmt.Println(client.Execute(command))
}

func (client *Client) EnsurePath(path string) {
	_, error := client.sshClient.Run("ls " + path)
	if error != nil {
		client.logger.Warn(error.Error())
		if len(client.TargetUser) > 0 {
			client.Execute("sudo -u " + client.TargetUser + " mkdir -p " + path)
		} else {
			client.Execute("mkdir -p " + path)
		}
	}
}

func (client *Client) PushFile(source string, target string) {
	client.logger.Info("Source for the upload: " + source)
	var adaptedTarget = target
	if len(client.TargetUser) > 0 && !strings.HasPrefix(target, "/") {
		adaptedTarget = "/home/" + client.TargetUser + "/" + target
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
	if len(client.TargetUser) > 0 {
		client.SetFileOwner(adaptedTarget, client.TargetUser)
	}
}

func (client *Client) UploadSftp(source string, target string) {
	err := client.sshClient.Upload(source, target)
	if err != nil {
		client.logger.Fatal("Error uploading: " + err.Error())
	}
}

func (client *Client) Upload(source string, target string) {
	clientConfig, _ := auth.PrivateKey(client.userName, client.keyPath, ssh.InsecureIgnoreHostKey())
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

func (client *Client) SetTargetUser(userName string) {
	client.TargetUser = userName
	client.ensureUserExists(userName)
}

func (client *Client) SetFileOwner(path string, userName string) {
	client.logger.Info("Ensure file " + path + " belongs to user " + userName + ".")
	result, _ := client.Execute("chown -R " + userName + ":" + userName + " " + path)
	if len(result) == 0 {
		client.logger.Info("Success")
	} else {
		client.logger.Info("Setting user permission")
	}
}

func (client *Client) EnsureExecutable(path string) {
	client.logger.Info("Ensure file is Executable ")
	if len(client.TargetUser) > 0 && !strings.HasPrefix(path, "/") {
		client.Execute("chmod +x " + "/home/" + client.TargetUser + "/" + path)
		client.logger.Info("Path: " + "/home/" + client.TargetUser + "/" + path)
		return
	}
	client.Execute("chmod +x " + path)
	client.logger.Info("Path: " + path)
}

func (client *Client) ensureUserExists(userName string) {
	client.logger.Info("Checking for user " + userName)
	result, _ := client.Execute("getent passwd " + userName)
	if len(result) == 0 {
		client.createUser(userName)
	} else {
		client.logger.Info("User " + userName + " exists.")
	}
}

func (client *Client) createUser(userName string) {
	client.logger.Info("User " + userName + " does not exist.")
	result, _ := client.Execute("useradd -m " + userName)
	if len(result) == 0 {
		client.logger.Info("User " + userName + " created successfully.")
	} else {
		client.logger.Fatal("Error during User creation")
	}
}

func (client *Client) TearDown() {
	defer func() {
		err := client.sshClient.Close()
		if err != nil {
			fmt.Printf("error closing ssh client %v", err)
		}
	}()
}
