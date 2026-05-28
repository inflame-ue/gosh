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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("gosh> ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		input := scanner.Text()
		if len(input) == 0 {
			fmt.Println()
			continue
		}

		commands, err := parser.ParseCommands(input)
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

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigc
		fmt.Println("\ninterrupt received...exiting...")
		os.Exit(1)
	}()
	repl()
}
