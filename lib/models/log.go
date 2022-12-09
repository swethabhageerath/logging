package models

type Log struct {
	AppName   string
	Username  string
	RequestId string
	Frames    []frame
	Message   string
	Params    string
	Details   string
}

type frame struct {
	FilePath   string
	Operation  string
	LineNumber int
}
