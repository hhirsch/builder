package traits

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
)

type FileSystem struct {
}

func NewFileSystem() *FileSystem {
	return &FileSystem{}
}

func (this *FileSystem) writeMapToTemporaryFileAsJson(path string, data map[string]string) (filePath string, err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	filePath = this.getTemporaryFilePath(string(jsonData))
	return filePath, this.writeMapToFileAsJson(path, data)
}

func (this *FileSystem) writeMapToFileAsJson(path string, data map[string]string) (err error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return this.WriteStringToFile(path, string(jsonData))
}

func (this *FileSystem) getTemporaryFilePath(content string) string {
	hash := md5.New()
	hash.Write([]byte(content))
	hashSum := hash.Sum(nil)
	return "/tmp/builder-" + hex.EncodeToString(hashSum)
}

func (this *FileSystem) WriteStringToTempFile(content string) (filePath string, err error) {
	filePath = this.getTemporaryFilePath(content)
	err = this.WriteStringToFile(filePath, content)
	return
}

func (this *FileSystem) WriteStringToFile(path string, content string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write([]byte(content))
	if err != nil {
		return err
	}
	return
}

func (this *FileSystem) getMapFromJsonFile(path string) (data map[string]string, err error) {
	var content string
	content, err = this.getStringFromFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(content), &data)
	return
}

func (this *FileSystem) getStringFromFile(path string) (content string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	var byteContent []byte
	byteContent, err = io.ReadAll(file)
	content = string(byteContent)
	if err != nil {
		return "", err
	}

	return
}
