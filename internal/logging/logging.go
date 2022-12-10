package logging

import (
	"io"
	i "libraries/logging/internal/logger"
	"libraries/logging/lib/constants"
	l "libraries/logging/lib/logger"
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

func (l Logging) getLogger() l.ILogger {
	switch l.loggerType {
	case constants.DEFAULT:
		return i.NewDefaultLogger()
	default:
		return i.NewDefaultLogger()
	}
}
