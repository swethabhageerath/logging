package logger

type ILogging interface {
	NewLogger() ILogger
}
