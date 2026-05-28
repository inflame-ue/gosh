package command

import (
	"os"
	"os/exec"
)

type RedirectState string

const (
	OutputToFile  RedirectState = ">"
	AppendToFile  RedirectState = ">>"
	InputFromFile RedirectState = "<"
)

type CommandRedirect struct {
	File  string
	State RedirectState
}

type Command struct {
	Name     string
	Args     []string
	Redirect *CommandRedirect
}

func NewCommand(name string, args []string) *Command {
	return &Command{
		Name:     name,
		Args:     args,
		Redirect: nil,
	}
}

func NewCommandWithRedirect(name string, args []string, redirect *CommandRedirect) *Command {
	return &Command{
		Name:     name,
		Args:     args,
		Redirect: redirect,
	}
}

func prepareFile(commandRedirect *CommandRedirect) (*os.File, error) {
	filepath := commandRedirect.File

	switch commandRedirect.State {
	case OutputToFile:
		file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return &os.File{}, err
		}
		return file, nil
	case AppendToFile:
		file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return &os.File{}, err
		}
		return file, nil
	case InputFromFile:
		file, err := os.OpenFile(filepath, os.O_RDONLY, 0)
		if err != nil {
			return &os.File{}, err
		}
		return file, nil
	default:
		return &os.File{}, nil
	}
}

func ExecuteCommands(cmds []*Command) error {
	var commands []*exec.Cmd
	for _, c := range cmds {
		commands = append(commands, exec.Command(c.Name, c.Args...))
	}

	for index := 0; index < len(commands)-1; index++ {
		pipe, err := commands[index].StdoutPipe()
		if err != nil {
			return err
		}
		commands[index+1].Stdin = pipe
	}

	var out *os.File
	lastCommand := cmds[len(cmds)-1]
	if lastCommand.Redirect != nil {
		file, err := prepareFile(lastCommand.Redirect)
		if err != nil {
			return err
		}
		defer file.Close()
		out = file
	} else {
		out = os.Stdout
	}
	commands[len(commands)-1].Stdout = out

	for _, cmd := range commands {
		err := cmd.Start()
		if err != nil {
			return err
		}
	}

	for _, cmd := range commands {
		err := cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Command) Execute() error {
	cmd := exec.Command(c.Name, c.Args...)

	if c.Redirect != nil {
		file, err := prepareFile(c.Redirect)
		if err != nil {
			return err
		}
		defer file.Close()

		if c.Redirect.State == OutputToFile || c.Redirect.State == AppendToFile {
			cmd.Stdin = os.Stdin
			cmd.Stdout = file
			cmd.Stderr = os.Stderr
		} else {
			cmd.Stdin = file
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
