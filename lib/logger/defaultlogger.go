package logger

import (
	"context"
	"libraries/logging/lib/constants"
	"libraries/logging/lib/models"
)

type DefaultLogger struct {
	log *models.Log
	//add writers
}

func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{
		log: new(models.Log),
	}
}

func (d *DefaultLogger) AddMessage(message string) {
	d.log.Message = message
}

func (d *DefaultLogger) AddParams(params string) {
	d.log.Params = params
}

func (d *DefaultLogger) AddDetails(details string) {
	d.log.Details = details
}

func (d *DefaultLogger) AddContext(context context.Context) {
}

func (d *DefaultLogger) AddFrames(frameDepth constants.FrameDepth) {

}
