package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inflame-ue/gosh/internal/builtins"
	"github.com/inflame-ue/gosh/internal/command"
	"github.com/inflame-ue/gosh/internal/parser"
)

func repl() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	input := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input <- scanner.Text()
		}
		input <- "\x04" // sentinal value to denote Ctrl-D signal
	}()

	for {
		fmt.Print("gosh> ")
		select {
		case <-sigc:
			fmt.Println()
			continue
		case line := <-input:
			if line == "\x04" {
				fmt.Print("\nexiting...goodbye...")
				os.Exit(0)
			}

			if line == "" {
				continue
			}
			
			commands, err := parser.ParseCommands(line)
			if err != nil {
				fmt.Printf("err: %v\n", err)
				continue
			}

			if len(commands) > 1 {
				err := command.ExecuteCommands(commands)
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}

				// TODO: the builtins do not support pipes?
				// i guess we skip to $PATH programs here anyway, so builtins are just for fancy demonstration purposes
				// that i am capable of implementing a subset of shell commands (why?)
				continue
			}

			command := commands[0]
			if handler, ok := builtins.BuiltinCommands[command.Name]; ok {
				handler(*command)
				continue
			}

			err = command.Execute()
			if err != nil {
				fmt.Printf("err: %v\n", err)
				continue
			}
		}
	}
}

func main() {
	repl()
}
