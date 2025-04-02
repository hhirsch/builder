package traits

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type FileSystem struct {
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (fileSystem *FileSystem) WriteMapToTemporaryFileAsJson(path string, data map[string]string) (filePath string, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	filePath = fileSystem.getTemporaryFilePath(string(jsonData))
	return filePath, fileSystem.writeMapToFileAsJson(path, data)
}

func (fileSystem *FileSystem) writeMapToFileAsJson(path string, data map[string]string) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return fileSystem.WriteStringToFile(path, string(jsonData))
}

func (fileSystem *FileSystem) getTemporaryFilePath(content string) string {
	hash := md5.New()
	hash.Write([]byte(content))
	hashSum := hash.Sum(nil)
	return "/tmp/builder-" + hex.EncodeToString(hashSum)
}

func (fileSystem *FileSystem) WriteStringToTempFile(content string) (filePath string, err error) {
	filePath = fileSystem.getTemporaryFilePath(content)
	err = fileSystem.WriteStringToFile(filePath, content)
	return
}

func (fileSystem *FileSystem) WriteStringToFile(path string, content string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}
	return
}

func (fileSystem *FileSystem) GetMapFromJsonFile(path string) (data map[string]string, err error) {
	var content string
	content, err = fileSystem.getStringFromFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		return
	}
	return
}

func (fileSystem *FileSystem) getStringFromFile(path string) (content string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file %v", err)
		}
	}()
	var byteContent []byte
	byteContent, err = io.ReadAll(file)
	content = string(byteContent)
	if err != nil {
		return "", err
	}

	return
}
