// needs to be replaced or sunset
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

func (hostRegistry *HostRegistry) getKeyPath(key string, hostName string) (string, error) {
	if hostName == "" {
		return "", fmt.Errorf("hostname is empty")
	}
	keyPath := "host." + hostName + "." + key

	return keyPath, nil
}

// todo get the hostname again
func (hostRegistry *HostRegistry) PromptEncryptedIfMissing(key string) (value string, err error) {
	var keyPath string
	keyPath, err = hostRegistry.getKeyPath(key, "")
	if err != nil {
		return "", err
	}
	_, err = hostRegistry.environment.GetRegistry().GetEncryptedString(keyPath)
	if err != nil {
		hostRegistry.environment.GetLogger().Infof("No host key for %s found in registry asking for user input.", key)
		inputField := huh.NewInput().
			Title("Enter a value for" + key).
			Value(&value)
		err = inputField.Run()
		if err != nil {
			return "", err
		}
		hostRegistry.environment.GetLogger().Info("Registering " + key)
		err = hostRegistry.environment.GetRegistry().RegisterEncrypted(keyPath, value)
		if err != nil {
			return "", err
		}
	}

	return value, err
}

func (hostRegistry *HostRegistry) PromptIfMissing(key string) (value string, err error) {
	var keyPath string
	keyPath, err = hostRegistry.getKeyPath(key)
	if err != nil {
		return "", err
	}
	_, err = hostRegistry.environment.GetRegistry().GetValue(keyPath)
	if err != nil {
		hostRegistry.environment.GetLogger().Infof("No host key for %s found in registry asking for user input.", key)
		input := huh.NewInput().
			Title("Enter a value for" + key).
			Value(&value)
		err := input.Run()
		if err != nil {
			return "", err
		}
		hostRegistry.environment.GetLogger().Info("Registering " + key)
		hostRegistry.environment.GetRegistry().Register(keyPath, value)
	}

	return value, nil
}
