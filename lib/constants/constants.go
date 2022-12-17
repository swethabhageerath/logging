package constants

type LogLevel int

const (
	ERROR LogLevel = iota
	DEBUG
)

func (l LogLevel) String() string {
	switch l {
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	default:
		return "DEBUG"
	}
}

const (
	KEY_LOGDIRECTORYPATH = "KEY_LOGDIRECTORYPATH"
)
