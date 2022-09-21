/*
This command_parser will read the file contents by iterating through each line, and returns them as <COMMAND> <ARGS...> format.

The command will be used to evaluate the task to be performed.

The args will be used to help the computation process.
*/
package command_parser

type CommandParser interface {
	// This will returns the list of commands to be processed by the business logic
	Commands() ([][]string, error)
}
