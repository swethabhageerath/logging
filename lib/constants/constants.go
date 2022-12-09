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
