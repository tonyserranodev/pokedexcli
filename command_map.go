package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	res, err := cfg.pokeAPIClient.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	for _, area := range res.Results {
		cfg.VisitedLocations[area.Name] = struct{}{}
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	res, err := cfg.pokeAPIClient.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = res.Next
	cfg.Previous = res.Previous
	for _, location := range res.Results {
		fmt.Println(location.Name)
	}
	return nil

}
