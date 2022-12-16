package models

type Log struct {
	Frames    string
	AppName   string
	User      string
	RequestId string
	LogLevel  string
	Message   string
	Params    string
	Details   string
}

type Frame struct {
	FilePath   string
	Operation  string
	LineNumber int
}
