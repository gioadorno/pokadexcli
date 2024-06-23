package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      description: "Prints the help menu",
      callback: callbackHelp,
    },
    "map": {
      name: "map",
      description: "Lists some location areas",
      callback: callbackMap,
    },
    "exit": {
      name: "exit",
      description: "Turns off the Pokedex",
      callback: callbackExit,
    },
  }
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokadex >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]

    availableCommands := getCommands()

    command, ok := availableCommands[commandName]
    if !ok {
      fmt.Println("invalid command")
      continue
    }

    command.callback(cfg)
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
