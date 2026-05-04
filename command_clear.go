package main

import (
	"fmt"
)

func commandClear(cfg *config, args ...string) error {
	fmt.Print("\033[H\033[2J")
	return nil
}
