package logger

import (
	"context"
	"fmt"
	"github.com/swethabhageerath/logging/lib/constants"
	"github.com/swethabhageerath/logging/lib/models"
	h "github.com/swethabhageerath/utilities/lib/utilities/helpers"
	"io"
	"runtime"
	"strings"
)

type DefaultLogger struct {
	log        *models.Log
	env        h.IEnvironmentHelper
	ctx        h.IContextHelper
	writers    []io.Writer
	marsheller h.IMarshallingHelper
}

func NewDefaultLogger(writers []io.Writer, env h.IEnvironmentHelper, ctx h.IContextHelper, marsheller h.IMarshallingHelper) *DefaultLogger {
	return &DefaultLogger{
		log:        &models.Log{},
		env:        env,
		ctx:        ctx,
		writers:    writers,
		marsheller: marsheller,
	}
}

func (d *DefaultLogger) AddAppName(appName string) ILogger {
	d.log.AppName = appName
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

func (d *DefaultLogger) AddContext(context context.Context) ILogger {
	requestIdKey := d.env.Get("KEY_REQUESTID")
	userKey := d.env.Get("KEY_USER")

	fmt.Println("RequestIdKey", requestIdKey, "UserKey", userKey)

	if requestIdKey == "" || userKey == "" {
		panic("RequestIdKey or UserKey not present in environment variables")
	}

	requestId, err := d.ctx.Get(context, requestIdKey)
	if err != nil {
		panic(err)
	}
	user, err := d.ctx.Get(context, userKey)
	if err != nil {
		panic(err)
	}

	d.log.RequestId = requestId
	d.log.User = user
	return d
}

func (d *DefaultLogger) AddFrames(frameDepth constants.FrameDepth) ILogger {
	if frameDepth == constants.CURRENT {
		d.log.Frames = append(d.log.Frames, d.getFrame())
	} else {
		d.log.Frames = d.getFrames()
	}
	return d
}

func (d *DefaultLogger) Log() {
	if d.log.LogLevel == "" {
		d.log.LogLevel = constants.DEBUG.String()
	}
	for _, i := range d.writers {
		s, err := h.MarshallingHelper{}.Marshall(d.log)
		fmt.Println(s)
		if err != nil {
			panic(err)
		}
		_, err = i.Write([]byte(s))
		if err != nil {
			panic(err)
		}
	}
}

func (d *DefaultLogger) getFrames() []models.Frame {
	ptrs := make([]uintptr, 10)
	numberOfEntries := runtime.Callers(0, ptrs)
	if numberOfEntries == 0 {
		return nil
	}

	ptrs = ptrs[:numberOfEntries]
	callerFrames := runtime.CallersFrames(ptrs)

	frames := make([]models.Frame, 0)

	for {
		frame, more := callerFrames.Next()

		if !strings.Contains(frame.File, "defaultlogger.go") && !strings.Contains(frame.File, "/runtime/") {
			frames = append(frames, models.Frame{
				FilePath:   frame.File,
				LineNumber: frame.Line,
				Operation:  frame.Func.Name(),
			})
		}

		if !more {
			break
		}
	}

	return frames
}

func (d *DefaultLogger) getFrame() models.Frame {
	ptr, file, line, _ := runtime.Caller(0)

	return models.Frame{
		FilePath:   file,
		Operation:  runtime.FuncForPC(ptr).Name(),
		LineNumber: line,
	}
}
