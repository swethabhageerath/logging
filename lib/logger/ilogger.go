package logger

import (
	"github.com/swethabhageerath/logging/lib/constants"
)

type ILogger interface {
	AddMessage(message string) ILogger
	AddParams(params string) ILogger
	AddDetails(details string) ILogger
	AddStackTrace(e error) ILogger
	AddLogLevel(logLevel constants.LogLevel) ILogger
	AddAppName(appName string) ILogger
	Log()
}
