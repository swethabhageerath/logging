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
