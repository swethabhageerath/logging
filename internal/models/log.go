package models

type Log struct {
	AppName   string
	User      string
	RequestId string
	Frames    []Frame
	Message   string
	Params    string
	Details   string
}

type Frame struct {
	FilePath   string
	Operation  string
	LineNumber int
}
