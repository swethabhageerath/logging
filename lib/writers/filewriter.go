package writers

import (
	"fmt"
	"os"

	h "github.com/swethabhageerath/utilities/lib/utilities/helpers"
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
	logFileParentDirectory, err := os.UserHomeDir()
	os.Setenv("KEY_LOGDIRECTORYPATH", "/logs")
	logFileDirectory := f.env.Get("KEY_LOGDIRECTORYPATH")
	fullLogDirectoryPath := fmt.Sprintf("%s%s", logFileParentDirectory, logFileDirectory)
	if err != nil {
		panic(err)
	}
	if logFileDirectory == "" {
		panic("Log File Directory path cannot be retrieved")
	}
	filePath, err := f.file.CreateFileWithCurrentDate(fullLogDirectoryPath)
	if err != nil {
		panic(err)
	}
	err = f.file.WriteFile(filePath, string(b)+"\n")
	if err != nil {
		panic(err)
	}
	return 0, nil
}
