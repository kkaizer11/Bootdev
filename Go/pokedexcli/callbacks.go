package main

import (
	"fmt"
	"os"
)

func command_help() error {
	available_commands := getCommands()
	fmt.Println("These are the available commands")
	for _, command := range available_commands {
		fmt.Printf("  %s: %s\n", command.name, command.description)
	}
	return nil
}

func command_exit() error {
	fmt.Println("Exiting the Pokedex")
	os.Exit(0)
	return nil
}
