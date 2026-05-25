package command

import "fmt"

type Command struct {
	Name string
	Args []string
	Flags map[string]string
}

func NewCommand(name string, args []string, flags map[string]string) *Command {
	return &Command{
		Name: name,
		Args: args,
		Flags: flags,
	}
}

func (c *Command) Execute() {
	fmt.Print("executing...\n")
}

