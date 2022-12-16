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

type FrameDepth int

const (
	CURRENT FrameDepth = iota
	ALL
)

func (f FrameDepth) String() string {
	switch f {
	case CURRENT:
		return "Current"
	case ALL:
		return "All"
	default:
		return "Current"
	}
}

type LoggerType int

const (
	DEFAULT LoggerType = iota
)

func (l LoggerType) String() string {
	switch l {
	case DEFAULT:
		return "Default"
	default:
		return "Default"
	}
}

type RequestKeys int

const (
	KEY_REQUESTID RequestKeys = iota
	KEY_USER
)

func (r RequestKeys) String() string {
	switch r {
	case KEY_REQUESTID:
		return "KEY_REQUESTID"
	case KEY_USER:
		return "KEY_USER"
	default:
		return ""
	}
}

type ErrorTypes int

const (
	ERR_REQUEST_CONTEXT_VALUE_KEYS_NOT_PROVIDED ErrorTypes = iota
	ERR_REQUEST_CONTEXT_VALUES_NOT_PROVIDED
)

func (e ErrorTypes) String() string {
	switch e {
	case ERR_REQUEST_CONTEXT_VALUES_NOT_PROVIDED:
		return "RequestId and/or User is not provided in context. They are mandatory"
	case ERR_REQUEST_CONTEXT_VALUE_KEYS_NOT_PROVIDED:
		return "RequestIdKey and/or UserKey is not provided in environment variables. They are mandatory"
	default:
		return ""
	}
}

type LoggerExcludesValues int

const (
	DEFAULT_LOGGER LoggerExcludesValues = iota
	RUNTIME
)

func (l LoggerExcludesValues) String() string {
	switch l {
	case DEFAULT_LOGGER:
		return "defaultlogger.go"
	case RUNTIME:
		return "/runtime/"
	default:
		return ""
	}
}
