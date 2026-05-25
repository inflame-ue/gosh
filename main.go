package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/inflame-ue/gosh/internal/parser"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("gosh> ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Printf("err: %v", err)
		}

		input := scanner.Text()
		if len(input) == 0 {
			fmt.Println()
			continue
		}
		
		command, err := parser.ParseCommand(input)
		if err != nil {
			log.Printf("err: %v", err)
		}
		
		err = command.Execute()
		if err != nil {
			log.Printf("err: %v", err)
		}

		fmt.Println()
	}
}

func main() {
	repl()
}
