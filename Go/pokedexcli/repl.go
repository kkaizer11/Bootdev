package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
	scanf := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanf.Scan()
		text := scanf.Text()
		cleanned := cleanInput(text)
		available_commands := getCommands()
		if len(cleanned) == 0 {
			continue
		}
		commandName := cleanned[0]
		command, ok := available_commands[commandName]
		if !ok {
			fmt.Println("Invalid Command, type \"help\" to display all availabe commands")
			continue
		}
		command.callback()
	}
}
