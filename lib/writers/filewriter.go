package writers

import (
	"fmt"
	h "github.com/swethabhageerath/utilities/lib/utilities/helpers"
	"os"
)

type FileWriter struct {
	env  h.IEnvironmentHelper
	file h.IFileHelper
}

func New(env h.IEnvironmentHelper, file h.IFileHelper) FileWriter {
	return FileWriter{
		env:  env,
		file: file,
	}
}

func (f FileWriter) Write(b []byte) (int, error) {
	logFileDirectory, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	if logFileDirectory == "" {
		panic("Log File Directory path cannot be retrieved")
	}
	logFileDirectory = fmt.Sprintf("%s/logs", logFileDirectory)
	filePath, err := f.file.CreateFileWithCurrentDate(logFileDirectory)
	if err != nil {
		panic(err)
	}
	err = f.file.WriteFile(filePath, string(b)+"\n")
	if err != nil {
		panic(err)
	}
	return 0, nil
}
