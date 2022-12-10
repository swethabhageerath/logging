package writers

import (
	"github.com/swethabhageerath/utilities/lib/utilities/environmenthelper"
	"github.com/swethabhageerath/utilities/lib/utilities/filehelpers"
)

type FileWriter struct {
	env  environmenthelper.IEnvironmentHelper
	file filehelpers.IFileHelper
}

func New(env environmenthelper.IEnvironmentHelper, file filehelpers.IFileHelper) FileWriter {
	return FileWriter{
		env:  env,
		file: file,
	}
}

func (f FileWriter) Write(b []byte) (int, error) {
	logFileDirectory := f.env.Get("KEY_LOGDIRECTORYPATH")
	if logFileDirectory == "" {
		panic("Log File Directory path is not specified in the environment variables")
	}
	filePath, err := f.file.CreateFileWithCurrentDate(logFileDirectory)
	if err != nil {
		panic(err)
	}
	err = f.file.WriteFile(filePath, string(b))
	if err != nil {
		panic(err)
	}
	return 0, nil
}
