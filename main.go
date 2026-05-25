package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/inflame-ue/gosh/internal/parser"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigc
			fmt.Println("\ninterrupt received...continuing...")
			os.Exit(1)
		}()

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

		command, err := parser.ParseCommand(input)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		err = command.Execute()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			continue
		}

		fmt.Println()
	}
}

func main() {
	repl()
}
