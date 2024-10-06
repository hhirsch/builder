package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Registry struct {
	data     map[string]string
	fileName string
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

func (this *Registry) GetBool(key string) (bool, error) {
	if this.data[key] == "true" {

		return true, nil
	}

	return false, nil
}

func (this *Registry) GetValue(key string) (string, error) {
	value, ok := this.data[key]
	if !ok {
		return "", errors.New("Value does not exist.")
	}
	return value, nil
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
