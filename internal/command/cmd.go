package command

import (
	"os/exec"
)

type Command struct {
	Name  string
	Args  []string
	Flags map[string]string
}

func NewCommand(name string, args []string, flags map[string]string) *Command {
	return &Command{
		Name:  name,
		Args:  args,
		Flags: flags,
	}
}

func (c *Command) Execute() ([]byte, error) {
	out, err := exec.Command(c.Name).Output()
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}
