package command_parser

type CommandParser interface {
	Commands() ([][]string, error)
}
