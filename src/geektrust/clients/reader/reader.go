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

// FileSystem - A FileSystem provides access to a file.
// The FileSystem interface is the minimum implementation required of the file.
type FileSystem interface {
	// Open the given file identified by name.
	// The returned file further can be used for other text related processing.
	//
	// Returns an error when failed to open the file.
	Open(name string) (fs.File, error)
}

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode O_RDONLY.
// If there is an error, it will be of type *PathError.
