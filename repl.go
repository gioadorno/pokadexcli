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
	callback    func(*config, ...string) error
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
      description: "Lists next page of areas",
      callback: callbackMap,
    },
    "mapb": {
      name: "mapb",
      description: "Lists previous page of areas",
      callback: callbackMapB,
    },
    "explore": {
      name: "explore {location_area}",
      description: "Lists the pokemon in a location area",
      callback: callbackExplore,
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
    args := []string{}
    if len(cleaned) > 1 {
      args = cleaned[1:]
    }

    availableCommands := getCommands()

    command, ok := availableCommands[commandName]
    if !ok {
      fmt.Println("invalid command")
      continue
    }

    err := command.callback(cfg, args...)
    if err != nil {
      fmt.Println(err)
    }
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
