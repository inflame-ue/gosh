package command

import (
	"os"
	"os/exec"
)

type Command struct {
	Name string
	Args []string
}

func NewCommand(name string, args []string) *Command {
	return &Command{
		Name: name,
		Args: args,
	}
}

func ExecuteCommands(cmds []*Command) error {
	var commands []*exec.Cmd
	for _, c := range cmds {
		commands = append(commands, exec.Command(c.Name, c.Args...))
	}

	for index := 0; index < len(commands) - 1; index++ {
		pipe, err := commands[index].StdoutPipe()
		if err != nil {
			return err
		}
		commands[index + 1].Stdin = pipe
	}
	commands[len(commands) - 1].Stdout = os.Stdout

	// TODO: ponder on whether there are any concurrency problems here?
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
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
