package models

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/swethabhageerath/logging/lib/constants"
)

type Options func(*log)

type log struct {
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

type Frame struct {
	FilePath   string
	Operation  string
	LineNumber int
}

func New(options ...Options) *log {
	log := new(log)
	for _, i := range options {
		i(log)
	}
	return log
}

func (l *log) Attach(observer io.Writer) (bool, error) {
	for _, o := range l.Observers {
		if o == observer {
			return false, errors.New("logging observer already attached")
		}
	}

	l.Observers = append(l.Observers, observer)
	return true, nil
}

func (l *log) Detach(observer io.Writer) (bool, error) {
	for i, o := range l.Observers {
		if o == observer {
			l.Observers = append(l.Observers[:i], l.Observers[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("logging observer not attached")
}

func (l *log) Notify() (bool, error) {
	for _, o := range l.Observers {
		_, err := o.Write([]byte(l.String()))
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func WithMandatoryFields(appName string, user string, loglevel constants.LogLevel) Options {
	return func(l *log) {
		if appName == "" {
			panic("appname is mandatory")
		}

		if user == "" {
			panic("user is mandatory")
		}

		l.LogLevel = loglevel.String()
		l.AppName = appName
		l.User = user
	}
}

func WithRequestId(requestId string) Options {
	return func(l *log) {
		if requestId == "" {
			l.RequestId = requestId
		}
	}
}

func WithDetails(details string) Options {
	return func(l *log) {
		if details == "" {
			l.Details = details
		}
	}
}

func WithParams(params string) Options {
	return func(l *log) {
		if params == "" {
			l.Params = params
		}
	}
}

func WithStackTrace(err error) Options {
	return func(l *log) {
		type stackTracer interface {
			StackTrace() errors.StackTrace
		}

		var stackBuilder strings.Builder

		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				stack := fmt.Sprintf("%+v ", f)

				if !strings.Contains(stack, "/runtime") {
					stackBuilder.WriteString(strings.Replace(stack, "\n\t", " ", -1))
				}
			}
		}

		l.Frames = stackBuilder.String()
	}
}

func (l *log) String() string {
	j, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return string(j)
}
