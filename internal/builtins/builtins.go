package builtins

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/inflame-ue/gosh/internal/command"
)

var BuiltinCommands = map[string]func(cmd command.Command){
	"cd":   cdHandler,
	"pwd":  pwdHandler,
	"echo": echoHandler,
	"exit": exitHandler,
	"help": helpHandler,
}

func writeToStd(out string, std *os.File) {
	writer := bufio.NewWriter(std)
	_, err := writer.WriteString(out + "\n")
	if err != nil {
		fmt.Printf("failed to write to std: %v", err)
	}
	writer.Flush()
}

func cdHandler(cmd command.Command) {
	home, err := os.UserHomeDir()
	if err != nil {
		writeToStd("err: failed to fetch home dir", os.Stderr)
	}

	if len(cmd.Args) == 0 {
		err = os.Chdir(home)
		if err != nil {
			writeToStd("err: failed to go to home dir", os.Stderr)
			return
		}
		return
	}

	path := cmd.Args[0]
	absPath := filepath.Join(home, path)
	err = os.Chdir(absPath)
	if err != nil {
		writeToStd("err: failed to change dir, invalid path?", os.Stderr)
		return
	}
}

func pwdHandler(cmd command.Command) {
	cwd, err := os.Getwd()
	if err != nil {
		writeToStd("err: failed to get current working dir", os.Stderr)
		return
	}

	writeToStd(cwd, os.Stdout)
}

func echoHandler(cmd command.Command) {
	message := strings.Join(cmd.Args, " ")
	writeToStd(message, os.Stdout)
}

func exitHandler(cmd command.Command) {
	if len(cmd.Args) == 0 {
		os.Exit(0)
	}

	code, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		writeToStd("err: failed to parse the exit code", os.Stderr)
		return
	}
	os.Exit(code)
}

func helpHandler(cmd command.Command) {
	writeToStd("Welcome to gosh! The list of built-in commands follows.", os.Stdout)
	writeToStd("Please note that flags for commands are not supported, they mimic basic behavior only.\n", os.Stdout)

	writeToStd("- help        : prints this help message", os.Stdout)
	writeToStd("- pwd         : prints the current working directory\n", os.Stdout)
	writeToStd("- exit [code] : exit the shell with exit code, 0 by default", os.Stdout)
	writeToStd("- echo [args] : echoes the provided args back to stdout", os.Stdout)
	writeToStd("- cd   [dir]  : changes the cwd to dir, if not provided changes to user home", os.Stdout)
}
