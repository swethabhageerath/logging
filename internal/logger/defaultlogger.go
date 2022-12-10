package logger

import (
	"context"
	c "github.com/swethabhageerath/utilities/lib/utilities/contexthelper"
	u "github.com/swethabhageerath/utilities/lib/utilities/environmenthelper"
	m "github.com/swethabhageerath/utilities/lib/utilities/marshallinghelper"
	"io"
	"libraries/logging/internal/models"
	"libraries/logging/lib/constants"
	"runtime"
)

type DefaultLogger struct {
	log        *models.Log
	env        u.IEnvironmentHelper
	ctx        c.IContextHelper
	writers    []io.Writer
	marsheller m.IMarshallingHelper
}

func NewDefaultLogger(writers []io.Writer, env u.IEnvironmentHelper, ctx c.IContextHelper, marsheller m.IMarshallingHelper) *DefaultLogger {
	return &DefaultLogger{
		env:        env,
		ctx:        ctx,
		writers:    writers,
		marsheller: marsheller,
	}
}

func (d *DefaultLogger) AddMessage(message string) {
	d.log.Message = message
}

func (d *DefaultLogger) AddParams(params string) {
	d.log.Params = params
}

func (d *DefaultLogger) AddDetails(details string) {
	d.log.Details = details
}

func (d *DefaultLogger) AddContext(context context.Context) {
	requestIdKey := d.env.Get("KEY_REQUESTID")
	userKey := d.env.Get("KEY_USER")

	if requestIdKey == "" || userKey == "" {
		panic("RequestIdKey or UserKey not present in environment variables")
	}

	requestId := d.ctx.Get(context, requestIdKey)
	user := d.ctx.Get(context, userKey)

	d.log.RequestId = requestId
	d.log.User = user
}

func (d *DefaultLogger) AddFrames(frameDepth constants.FrameDepth) {
	if frameDepth == constants.CURRENT {
		d.log.Frames = append(d.log.Frames, d.getFrame())
	} else {
		d.log.Frames = d.getFrames()
	}
}

func (d *DefaultLogger) Log() {
	for _, i := range d.writers {
		s, err := m.MarshallingHelper{}.Marshall(d.log)
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

	ptrs = ptrs[:numberOfEntries]
	callerFrames := runtime.CallersFrames(ptrs)

	frames := make([]models.Frame, len(ptrs))

	for {
		frame, more := callerFrames.Next()

		if !more {
			break
		}

		frames = append(frames, models.Frame{
			FilePath:   frame.File,
			LineNumber: frame.Line,
			Operation:  frame.Func.Name(),
		})
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
