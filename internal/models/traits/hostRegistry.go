package traits

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/hhirsch/builder/internal/models"
)

type HostRegistry struct {
	environment *models.Environment
}

func NewHostRegistry(environment *models.Environment) *HostRegistry {
	return &HostRegistry{environment: environment}
}

func (this *HostRegistry) getKeyPath(key string) (string, error) {
	var hostName = this.environment.Client.GetHost()
	if hostName == "" {
		return "", fmt.Errorf("Hostname is empty.")
	}
	keyPath := "host." + this.environment.Client.GetHost() + "." + key

	return keyPath, nil
}

func (this *HostRegistry) PromptEncryptedIfMissing(key string) (value string, err error) {
	var keyPath string
	keyPath, err = this.getKeyPath(key)
	if err != nil {
		return "", err
	}
	_, err = this.environment.GetRegistry().GetEncryptedString(keyPath)
	if err != nil {
		this.environment.GetLogger().Infof("No host key for %s found in registry asking for user input.", key)
		inputField := huh.NewInput().
			Title("Enter a value for" + key).
			Value(&value)
		err = inputField.Run()
		if err != nil {
			return "", err
		}
		this.environment.GetLogger().Info("Registering " + key)
		err = this.environment.GetRegistry().RegisterEncrypted(keyPath, value)
		if err != nil {
			return "", err
		}
	}

	return value, err
}

func (this *HostRegistry) PromptIfMissing(key string) (value string, err error) {
	var keyPath string
	keyPath, err = this.getKeyPath(key)
	if err != nil {
		return "", err
	}
	_, err = this.environment.GetRegistry().GetValue(keyPath)
	if err != nil {
		this.environment.GetLogger().Infof("No host key for %s found in registry asking for user input.", key)
		input := huh.NewInput().
			Title("Enter a value for" + key).
			Value(&value)
		err := input.Run()
		if err != nil {
			return "", err
		}
		this.environment.GetLogger().Info("Registering " + key)
		this.environment.GetRegistry().Register(keyPath, value)
	}

	return value, nil
}
