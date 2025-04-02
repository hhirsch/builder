package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Registry struct {
	data       map[string]string
	fileName   string
	encryption *Encryption
}

func NewRegistry(fileName string) *Registry {
	data := make(map[string]string)
	return &Registry{
		data:     data,
		fileName: fileName,
	}
}

func (registry *Registry) Register(key string, value string) {
	registry.data[key] = value
}

func (registry *Registry) EnableRsa(encryption Encryption) {
	registry.encryption = &encryption
}

func (registry *Registry) RegisterEncrypted(key string, value string) (err error) {
	if registry.encryption == nil {
		err = errors.New("encryption is disabled")
	}
	registry.data[key], _ = registry.encryption.Encrypt(value)
	return
}

func (registry *Registry) GetEncryptedString(key string) (string, error) {
	if registry.encryption == nil {
		return "", errors.New("encryption is disabled")
	}

	value, err := registry.encryption.Decrypt(registry.data[key])
	if err == nil {
		return "", errors.New("value does not exist")
	}
	return value, nil
}

func (registry *Registry) GetBool(key string) (value bool, err error) {
	stringValue, exists := registry.data[key]
	if !exists {
		err = fmt.Errorf("key %s not found", key)
		return
	}

	switch strings.ToLower(stringValue) {
	case "true":
		value = true
	case "false":
		value = false
	default:
		err = fmt.Errorf("value for key %s is not a valid boolean string", key)
	}
	return
}

func (registry *Registry) GetValue(key string) (string, error) {
	value, exists := registry.data[key]
	if !exists {
		return "", errors.New("value does not exist")
	}
	return value, nil
}

func (registry *Registry) EraseValue(key string) error {
	_, exists := registry.data[key]
	if !exists {
		return errors.New("no value to remove")
	}
	delete(registry.data, key)
	return nil
}

func (registry *Registry) Save() (err error) {
	file, err := os.Create(registry.fileName)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()

	jsonData, err := json.Marshal(registry.data)
	if err != nil {
		return
	}

	_, err = file.Write(jsonData)
	return
}

func (registry *Registry) Load() (err error) {
	file, err := os.Open(registry.fileName)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(jsonData, &registry.data)
	if err != nil {
		return
	}

	return
}

func (registry *Registry) GetRegistryDump() map[string]string {

	return registry.data
}
