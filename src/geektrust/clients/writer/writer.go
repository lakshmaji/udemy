/*
Package writer - handles the responsibility of handling output operations.
*/
package writer

// BaseWriter - You can implement this writer interface to handle output representations for REST API's using http library or web framework.
//
// Current implementation contains code for handle STDOUT operations.
// Handles responsibility of writing to given Writer.
type BaseWriter interface {
	// Write output with formatter string and return carriage.
	WriteLn(format string, content ...interface{})
	// Write output.
	// panic and stop execution control flow when `panic` is set to true.
	// Exit the program execution when `panic` is set to false.
	WriteError(content interface{})
}
