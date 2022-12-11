package logger

import (
	"github.com/swethabhageerath/logging/lib/constants"
	h "github.com/swethabhageerath/utilities/lib/utilities/helpers"
	"io"
)

type Logging struct {
	writers    []io.Writer
	loggerType constants.LoggerType
	env        h.EnvironmentHelper
	ctx        h.ContextHelper
	mars       h.MarshallingHelper
}

func New(writers []io.Writer, loggerType constants.LoggerType) Logging {
	return Logging{
		writers:    writers,
		loggerType: loggerType,
		ctx:        h.ContextHelper{},
		env:        h.EnvironmentHelper{},
		mars:       h.MarshallingHelper{},
	}
}

func (l Logging) NewLogger() ILogger {
	switch l.loggerType {
	case constants.DEFAULT:
		return NewDefaultLogger(l.writers, l.env, l.ctx, l.mars)
	default:
		return NewDefaultLogger(l.writers, l.env, l.ctx, l.mars)
	}
}
