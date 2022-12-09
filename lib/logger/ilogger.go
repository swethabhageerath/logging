package logger

import (
	"context"
	"libraries/logging/lib/constants"
)

type ILogger interface {
	AddMessage(message string)
	AddParams(params string)
	AddDetails(details string)
	AddFrames(frameDepth constants.FrameDepth)
	AddContext(context context.Context)
}
