/*
Package cmd - will read the io.Reader contents by iterating through each line delimited by carriage return, and returns them as <COMMAND> <ARGS...> format.

The command will be used to evaluate the task to be performed. Returns error, when the io.Reader implementation fails to provide file contents.
*/
package cmd

// CommandParser ...
type CommandParser interface {
	// This will returns the list of commands to be processed by the business logic
	Commands() ([][]string, error)
}
