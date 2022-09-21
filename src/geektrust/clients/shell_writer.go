package clients

import (
	"fmt"
	"io"
	"os"
)

type Options struct {
	panic bool
}

var DefaultOptions *Options = &Options{}

type shellClient struct {
	w       io.Writer
	options *Options
}

// Handles responsibility of writing to **STDOUT**
func NewShellWriter(w io.Writer, options *Options) BaseWriter {
	return &shellClient{
		w:       w,
		options: options,
	}
}

func (s *shellClient) WriteLn(format string, content ...interface{}) {
	fmt.Fprintf(s.w, fmt.Sprintln(format), content...)
}

func (s *shellClient) WriteError(content interface{}) {
	fmt.Fprint(s.w, content)
	if s.options.panic {
		panic(content)
	}
	os.Exit(1)
}
