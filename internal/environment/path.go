package environment

import (
	"errors"
	"os"
	"os/user"
)

func getHomePath() (string, error) {
	currentUser, error := user.Current()
	if error != nil {
		return "", error
	}

	return currentUser.HomeDir, nil
}

func GetLogFileName() string {
	return "builder.log"
}

func GetProjectPath() string {
	return ".builder"
}

func GetProjectCommandsPath() string {
	return GetProjectPath() + "/commands/"
}

func GetGlobalRegistryPath() (string, error) {
	homePath, error := GetBuilderHomePath()
	if error != nil {
		return "", error
	}
	return homePath + "/builderGlobal.reg", nil
}

func GetKeyPath() (string, error) {
	homePath, error := getHomePath()
	if error != nil {
		return "", error
	}
	return homePath + "/.ssh/id_rsa", nil
}

func GetBuilderHomePath() (string, error) {
	homePath, error := getHomePath()
	if error != nil {
		return "", error
	}
	absoluteHomePath := homePath + "/" + GetProjectPath()
	_, error = os.Stat(absoluteHomePath)
	if errors.Is(error, os.ErrNotExist) {
		error := os.Mkdir(absoluteHomePath, os.ModePerm)
		if error != nil {
			return "", error
		}
	}

	return absoluteHomePath, nil
}

func GetLogFilePath() (string, error) {
	builderHomePath, error := GetBuilderHomePath()
	if error != nil {
		return "", error
	}
	return builderHomePath + "/" + GetLogFileName(), nil
}
