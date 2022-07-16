package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	EXIT_COMMAND   = ".exit"
	UNKNOW_COMMAND = "unknow"
)

func main() {
	for {
		printPrompt()
		command := readCommand()
		if command == EXIT_COMMAND {
			break
		}
		fmt.Printf("Unknow Command: %s\n", command)
	}

	fmt.Println("Exit izidb")
}

func printPrompt() {
	fmt.Print("izidb >> ")
}

func readCommand() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	command := scanner.Text()
	return command
}
