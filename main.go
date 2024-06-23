package main

import (
	"bufio"
	"fmt"
	"os"
)

// Define cliCommand
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Define the callback functions
func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
    Usage:

    help: Displays a help message
    exit: Exit the Pokedex`)
	return nil
}

func commandExit() error {
  fmt.Println("Exiting Pokedex...")
  os.Exit(0)
  return nil
}

// Initialize commands
func getCommands() map[string]cliCommand {
  return map[string]cliCommand{
    "help": {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
    "exit": {
      name: "exit",
      description: "Exit the Pokedex",
      callback: commandExit,
    },
  }
}

// Main REPL loop
func main() {
  commands := getCommands()
  scanner := bufio.NewScanner(os.Stdin)

  for {
    // Print prompt
    fmt.Print("Pokedex > ")
    if !scanner.Scan() {
      break
    }
    input := scanner.Text()

    // Check if the command exists in the  map
    command, exists := commands[input]
    if exists {
      // Execute the corresponding callback function
      err := command.callback()
      if err != nil {
        fmt.
          Println("Error:", err)
      }
    } else {
      fmt.Println("Unknown command:", input)
    }
  }
}
