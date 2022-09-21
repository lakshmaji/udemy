package clients

// Handles responsibility of writing to given Writer
type BaseWriter interface {
	Write(format string, content ...interface{})
	WriteLn(format string, content ...interface{})
	WriteError(interface{})
}
