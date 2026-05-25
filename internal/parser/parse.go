package parser

import (
	"strings"

	"github.com/inflame-ue/gosh/internal/command"
)

func ParseCommand(input string) (*command.Command, error) {
	commandParts := strings.Fields(input)
	commandName, commandArgs := commandParts[0], commandParts[1:]

	// TODO: for now we only care about the commandName, args and flags come later
	command := command.NewCommand(commandName, commandArgs) 

	return command, nil
}
