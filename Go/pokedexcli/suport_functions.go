package main

import "strings"

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    command_help,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    command_exit,
		},
	}
}
