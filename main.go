package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	EXIT_COMMAND   = ".exit"
	UNKNOW_COMMAND = "unknow"
)

func main() {
	command := "Init"
	for command != EXIT_COMMAND {
		printPrompt()
		command = readCommand()
		if command == "" {
			continue
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

	return strings.Trim(command, " ")
}
