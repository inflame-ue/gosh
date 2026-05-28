package parser

import (
	"errors"
	"os"
	"path"
	"strings"

	"github.com/inflame-ue/gosh/internal/command"
)

func parseFile(cmdString string, sep string) (string, error) {
	filename := strings.Split(cmdString, sep)[1]

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(cwd, filename), nil
}

func parseCommandRedirect(cmdString string) (*command.CommandRedirect, error) {
	var redirect command.CommandRedirect
	if strings.Contains(cmdString, ">>") {
		filepath, err := parseFile(cmdString, ">>")
		if err != nil {
			return nil, err
		}
		redirect = command.CommandRedirect{
			State: command.AppendToFile,
			File: filepath,
		}
	} else if strings.Contains(cmdString, ">") {
		filepath, err := parseFile(cmdString, ">")
		if err != nil {
			return nil, err
		}
		redirect = command.CommandRedirect{
			State: command.OutputToFile,
			File: filepath,
		}
	} else if strings.Contains(cmdString, "<") {
		filepath, err := parseFile(cmdString, "<")
		if err != nil {
			return nil, err
		}
		redirect = command.CommandRedirect{
			State: command.InputFromFile,
			File: filepath,
		}
	} else {
		return nil, errors.New("err: no valid redirect characters recognized")
	}
	return &redirect, nil
}

func ParseCommands(input string) ([]*command.Command, error) {
	var commands []*command.Command 
	
	commandStrings := strings.SplitSeq(input, "|")
	for commandString := range commandStrings {
		redirect, err := parseCommandRedirect(commandString)
		if err != nil {
			return nil, err
		}
		
		commandParts := strings.Fields(strings.Split(commandString, string(redirect.State))[0])
		commandName, commandArgs := commandParts[0], commandParts[1:]

		if redirect == nil {
			commands = append(commands, command.NewCommand(commandName, commandArgs))
		} else {
			commands = append(commands, command.NewCommandWithRedirect(commandName, commandArgs, redirect))
		}
	}
	
	return commands, nil
}
