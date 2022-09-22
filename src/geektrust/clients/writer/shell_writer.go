package writer

import (
	"fmt"
	"io"
	"os"
)

// Options is a configuration settings to customize behavior while writing output to shell STDOUT.
type Options struct {
	// panic is a setting to panic when an error encountered. It stops the program execution.
	//
	// default value is false
	//
	// When an error encountered the program will exit immediately without panic when this setting is not enabled.
	Panic bool
}

// Default settings
var DefaultOptions *Options = &Options{}

// Shell writer client
type shellClient struct {
	w       io.Writer
	options *Options
}

// This reader is shell based, which means it can interact with local file system to write to the shell.
//
// It expects Options.panic parameter, which will be used to define custom behavior when errors are encountered. This feature will be useful in identifying the risks early in development environment.
func New(w io.Writer, options *Options) BaseWriter {
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
	if s.options.Panic {
		panic(content)
	}
	os.Exit(1)
}
