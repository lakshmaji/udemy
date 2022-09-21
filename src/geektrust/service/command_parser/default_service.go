package command_parser

type service struct {
}

// Command parser for shell client input
func NewShellCommandParser() CommandParser {
	return &service{}
}
