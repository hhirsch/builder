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

func (this *Registry) Register(key string, value string) {
	this.data[key] = value
}

func (this *Registry) EnableRsa(encryption Encryption) {
	this.encryption = &encryption
}

func (this *Registry) RegisterEncrypted(key string, value string) (err error) {
	if this.encryption == nil {
		err = errors.New("encryption is disabled")
	}
	this.data[key], _ = this.encryption.Encrypt(value)
	return
}

func (this *Registry) GetEncryptedString(key string) (string, error) {
	if this.encryption == nil {
		return "", errors.New("encryption is disabled")
	}

	value, err := this.encryption.Decrypt(this.data[key])
	if err == nil {
		return "", errors.New("value does not exist")
	}
	return value, nil
}

func (this *Registry) GetBool(key string) (value bool, err error) {
	stringValue, exists := this.data[key]
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

func (this *Registry) GetValue(key string) (string, error) {
	value, exists := this.data[key]
	if !exists {
		return "", errors.New("value does not exist")
	}
	return value, nil
}

func (this *Registry) EraseValue(key string) error {
	_, exists := this.data[key]
	if !exists {
		return errors.New("no value to remove")
	}
	delete(this.data, key)
	return nil
}

func (this *Registry) Save() (err error) {
	file, err := os.Create(this.fileName)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()

	jsonData, err := json.Marshal(this.data)
	if err != nil {
		return
	}

	_, err = file.Write(jsonData)
	return
}

func (this *Registry) Load() (err error) {
	file, err := os.Open(this.fileName)
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

	err = json.Unmarshal(jsonData, &this.data)
	if err != nil {
		return
	}

	return
}

func (this *Registry) GetRegistryDump() map[string]string {

	return this.data
}
