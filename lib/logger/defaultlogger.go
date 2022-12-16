package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/swethabhageerath/logging/lib/constants"
	"github.com/swethabhageerath/logging/lib/models"
)

type DefaultLogger struct {
	log     *models.Log
	writers []io.Writer
}

func NewDefaultLogger(writers []io.Writer) *DefaultLogger {
	return &DefaultLogger{
		log:     &models.Log{},
		writers: writers,
	}
}

func (d *DefaultLogger) AddAppName(appName string) ILogger {
	d.log.AppName = appName
	return d
}

func (d *DefaultLogger) AddRequestId(requestId string) ILogger {
	if requestId == "" {
		panic("requestId is mandatory")
	}
	d.log.RequestId = requestId
	return d
}

func (d *DefaultLogger) AddUser(user string) ILogger {
	if user == "" {
		panic("user is mandatory")
	}
	d.log.User = user
	return d
}

func (d *DefaultLogger) AddMessage(message string) ILogger {
	d.log.Message = message
	return d
}

func (d *DefaultLogger) AddParams(params string) ILogger {
	d.log.Params = params
	return d
}

func (d *DefaultLogger) AddDetails(details string) ILogger {
	d.log.Details = details
	return d
}

func (d *DefaultLogger) AddLogLevel(logLevel constants.LogLevel) ILogger {
	l := logLevel.String()
	if l == "" {
		d.log.LogLevel = constants.DEBUG.String()
	} else {
		d.log.LogLevel = logLevel.String()
	}

	return d
}

func (d *DefaultLogger) AddStackTrace(e error) ILogger {
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

	d.log.Frames = stackBuilder.String()
	return d
}

func (d *DefaultLogger) Log() {
	if d.log.LogLevel == "" {
		d.log.LogLevel = constants.DEBUG.String()
	}
	for _, i := range d.writers {
		j, err := json.Marshal(d.log)
		if err != nil {
			panic(err)
		}

		_, err = i.Write([]byte(j))
		if err != nil {
			panic(err)
		}
	}
}

func (d *DefaultLogger) getEnvVariable(k constants.RequestKeys) string {
	return os.Getenv(k.String())
}

func (d *DefaultLogger) getContextVariable(ctx context.Context) (requestId string, user string) {
	requestIdKey := d.getEnvVariable(constants.KEY_REQUESTID)
	userKey := d.getEnvVariable(constants.KEY_USER)

	if requestIdKey == "" || userKey == "" {
		panic(constants.ERR_REQUEST_CONTEXT_VALUE_KEYS_NOT_PROVIDED)
	}

	requestId = ctx.Value(requestIdKey).(string)
	if requestId == "" {
		panic(constants.ERR_REQUEST_CONTEXT_VALUES_NOT_PROVIDED)
	}

	user = ctx.Value(userKey).(string)
	if user == "" {
		panic(constants.ERR_REQUEST_CONTEXT_VALUES_NOT_PROVIDED)
	}

	return
}
