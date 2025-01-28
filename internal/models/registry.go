package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
		err = errors.New("Encryption is disabled.")
	}
	this.data[key] = value
	return
}

func (this *Registry) GetEncryptedString(key string) (value string, err error) {
	value = ""
	if this.encryption == nil {
		err = errors.New("Encryption is disabled.")
	}

	value, ok := this.data[key]
	if !ok {
		err = errors.New("Value does not exist.")
	}
	return
}

func (this *Registry) GetBool(key string) (value bool, err error) {
	stringValue, exists := this.data[key]
	if !exists {
		err = fmt.Errorf("Key %s not found.", key)
		return
	}

	switch strings.ToLower(stringValue) {
	case "true":
		value = true
	case "false":
		value = false
	default:
		err = fmt.Errorf("Value for key %s is not a valid boolean string.", key)
	}
	return
}

func (this *Registry) GetValue(key string) (string, error) {
	value, exists := this.data[key]
	if !exists {
		return "", errors.New("Value does not exist.")
	}
	return value, nil
}

func (this *Registry) EraseValue(key string) error {
	_, exists := this.data[key]
	if !exists {
		return errors.New("No value to remove.")
	}
	delete(this.data, key)
	return nil
}

func (this *Registry) EncryptionTest() error {
	this.Register("encryptionTest", "test")
	testValue, err := this.GetValue("encryptionTest")
	if testValue != "test" && err == nil {
		return errors.New("Regular value retrieval failed.")
	}

	this.RegisterEncrypted("encryptionTest", "encrypted test")
	testValueUnencrypted, err := this.GetValue("encryptionTest")
	if testValueUnencrypted == "encrpyted test" {
		return errors.New("Unencrypted value retrieved after writing it encrypted.")
	}

	testValueEncrypted, err := this.GetEncryptedString("encryptionTest")
	if testValueEncrypted != "encrypted test" {
		return errors.New("Value was not encrypted.")
	}

	return nil
}

func (this *Registry) Save() (err error) {
	file, err := os.Create(this.fileName)
	if err != nil {
		return
	}
	defer file.Close()

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
	defer file.Close()

	jsonData, err := ioutil.ReadAll(file)
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
