package writers

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/swethabhageerath/logging/lib/constants"
)

type FileWriter struct{}

func (f FileWriter) Write(data []byte) (int, error) {
	logFileParentDirectory := f.getUserHomeDirectory()

	logFileDirectory := f.getLogDirectoryPath()

	fullLogDirectoryPath := path.Join(logFileParentDirectory, logFileDirectory)

	filePath := f.createFileWithCurrentDate(fullLogDirectoryPath)

	f.writeFile(filePath, data)

	return 0, nil
}

func (f FileWriter) writeFile(filePath string, data []byte) {
	fi, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(fi)
	}

	_, err = fi.Write(data)
	if err != nil {
		panic(err)
	}
}

func (f FileWriter) createFileWithCurrentDate(directoryPath string) string {
	fileNameWithCurrentDate := fmt.Sprintf("%s.txt", time.Now().Format("2006-01-02"))

	fullFilePath := path.Join(directoryPath, fileNameWithCurrentDate)

	if _, err := os.Stat(fullFilePath); err == nil {
		return fullFilePath
	}

	_, err := os.Create(fullFilePath)
	if err != nil {
		panic(err)
	}

	return fullFilePath
}

func (f FileWriter) getUserHomeDirectory() string {
	logFileParentDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return logFileParentDirectory
}

func (f FileWriter) getLogDirectoryPath() string {
	logFileDirectory := os.Getenv(constants.KEY_LOGDIRECTORYPATH)
	if logFileDirectory == "" {
		panic("logfiledirectory path is mandatory")
	}

	return logFileDirectory
}
