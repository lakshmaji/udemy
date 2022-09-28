/*
Package reader handles the responsibility of handling input operations.
*/
package reader

import (
	"io"
	"io/fs"
)

// BaseReader - You can implement this reader interface to handle inputs from REST API's using  http library or web framework.
//
// Current implementation contains code for handle STDIN operations.
type BaseReader interface {
	// Read the file by iterating over each line,
	//
	// and returns all the lines as an array of strings in the end.
	ParseFileName() (string, error)
	ParseFileContent(name string) (io.Reader, error)
	ParseFileLines(file io.Reader) ([]string, error)
}

// FileSystem - A FileSystem provides access to a file for Geekdemy.
type FileSystem interface {
	// Open the given file identified by name.
	Open(name string) (fs.File, error)
}
