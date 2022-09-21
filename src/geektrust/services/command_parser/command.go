package command_parser

type CommandParser interface {
	// This will returns the list of commands to be processed by the business logic
	Commands() ([][]string, error)
}
