package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokadex >")

		scanner.Scan()
		text := scanner.Text()

    cleaned := cleanInput(text)

		fmt.Print("echoing: ", cleaned)
	}

}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
