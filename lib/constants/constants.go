package constants

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
