package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("gosh> ")
		scanner.Scan()

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		command := scanner.Text()

		fmt.Printf("%v\n", command)
		fmt.Println()
	}
}

func main() {
	repl()
}
