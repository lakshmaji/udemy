/*
reader handles the responsibility of handling input operations.
*/
package reader

// You can implement this reader interface to handle inputs from REST API's using  http library or web framework.
//
// Current implementation contains code for handle STDIN operations.
type BaseReader interface {
	// Read the file by iterating over each line,
	//
	// and returns all the lines as an array of strings in the end.
	FileInput() ([]string, error)
}
