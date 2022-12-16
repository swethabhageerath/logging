package logger

import (
	"context"

	"github.com/swethabhageerath/logging/lib/constants"
)

type ILogger interface {
	AddMessage(message string) ILogger
	AddParams(params string) ILogger
	AddDetails(details string) ILogger
	AddStackTrace(e error) ILogger
	AddContext(context context.Context) ILogger
	AddLogLevel(logLevel constants.LogLevel) ILogger
	AddAppName(appName string) ILogger
	Log()
}
