package logging

import "io"

type Logging struct {
	writers []io.Writer
}

func New(writers []io.Writer) *Logging {
	return &Logging{
		writers: writers,
	}
}

func (l Logging) NewLogger() {

}
