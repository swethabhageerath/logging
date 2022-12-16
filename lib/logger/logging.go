package logger

import (
	"io"

	"github.com/swethabhageerath/logging/lib/constants"
)

type Logging struct {
	writers    []io.Writer
	loggerType constants.LoggerType
}

func New(writers []io.Writer, loggerType constants.LoggerType) Logging {
	return Logging{
		writers:    writers,
		loggerType: loggerType,
	}
}

func (l Logging) NewLogger() ILogger {
	switch l.loggerType {
	case constants.DEFAULT:
		return NewDefaultLogger(l.writers)
	default:
		return NewDefaultLogger(l.writers)
	}
}
