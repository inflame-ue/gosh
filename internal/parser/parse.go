package parser

import (
	"strings"

	"github.com/inflame-ue/gosh/internal/command"
)

func ParseCommands(input string) ([]*command.Command, error) {
	var commands []*command.Command 
	
	commandStrings := strings.SplitSeq(input, "|")
	for commandString := range commandStrings {
		commandParts := strings.Fields(commandString)
		commandName, commandArgs := commandParts[0], commandParts[1:]

		commands = append(commands, command.NewCommand(commandName, commandArgs))
	}
	
	return commands, nil
}
