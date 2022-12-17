package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/swethabhageerath/logging/lib/constants"
)

type Log struct {
	Frames    string `json:"Frames"`
	AppName   string `json:"AppName"`
	User      string `json:"User"`
	RequestId string `json:"RequestId"`
	LogLevel  string `json:"LogLevel"`
	Message   string `json:"Message"`
	Params    string `json:"Params"`
	Details   string `json:"Details"`

	Observers []io.Writer `json:"_"`
}

func (l *Log) Attach(observer io.Writer) (bool, error) {
	for _, o := range l.Observers {
		if o == observer {
			return false, errors.New("logging observer already attached")
		}
	}

	l.Observers = append(l.Observers, observer)
	return true, nil
}

func (l *Log) Detach(observer io.Writer) (bool, error) {
	for i, o := range l.Observers {
		if o == observer {
			l.Observers = append(l.Observers[:i], l.Observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("logging observer not attached")
}

func (l *Log) Notify() (bool, error) {
	for _, o := range l.Observers {
		_, err := o.Write([]byte(l.String()))
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (l *Log) AddAppName(appName string) *Log {
	if appName == "" {
		panic("appname is mandatory")
	}
	l.AppName = appName
	return l
}

func (l *Log) AddUser(user string) *Log {
	if user == "" {
		panic("user is mandatory")
	}
	l.User = user
	return l
}

func (l *Log) AddLogLevel(logLevel constants.LogLevel) *Log {
	l.LogLevel = logLevel.String()
	return l
}

func (l *Log) AddRequestId(requestId string) *Log {
	if requestId == "" {
		panic("requestid is mandatory")
	}
	l.RequestId = requestId
	return l
}

func (l *Log) AddDetails(details string) *Log {
	l.Details = details
	return l
}

func (l *Log) AddParams(params string) *Log {
	l.Params = params
	return l
}

func (l *Log) AddStackTrace(e error) *Log {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var stackBuilder strings.Builder

	if err, ok := e.(stackTracer); ok {
		for _, f := range err.StackTrace() {
			stack := fmt.Sprintf("%+v ", f)

			if !strings.Contains(stack, "/runtime") {
				stackBuilder.WriteString(strings.Replace(stack, "\n\t", " ", -1))
			}
		}
	}

	l.Frames = stackBuilder.String()
	return l
}

func (l *Log) String() string {
	j, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return string(j)
}

type Frame struct {
	FilePath   string
	Operation  string
	LineNumber int
}
