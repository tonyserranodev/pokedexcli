package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cfg := &Config{}
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands(cfg)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, ok := commands[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
