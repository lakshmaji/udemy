package writer

// Handles responsibility of writing to given Writer
type BaseWriter interface {
	WriteLn(format string, content ...interface{})
	WriteError(interface{})
}
