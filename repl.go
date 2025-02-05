package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		words := cleanInput(scanner.Text())

		if len(words) > 0 {
			fmt.Println("Your command was:", words[0])
		}

		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading stdin: %v", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
