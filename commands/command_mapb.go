package commands

import (
	"errors"
	"fmt"
)

func commandMapb(cfg *Config) error {
	// if the list was not previously started
	if cfg.prevCommand != "map" && cfg.offset == 0 {
		return errors.New("Use 'map' to start listing.")
	}

	// if list is 20 means that only first page was shown
	// TODO: make another field in 'Config' struct to keep track of more lines added
	// or even keep track of the total lines added and then 'clean' the console based on that
	if cfg.offset <= 20 {
		return errors.New("You are already at page 0!")
	}

	// the previous command has to be either map or mapb, so clear it
	for i := 0; i <= (24 + 1); i++ {
		fmt.Print("\033[A\033[2K")
	}

	// 1 step is '20', and both map and mapb advances 1 step at the end of the command
	// we need to take 2 steps back to get the previous areas to the last map
	cfg.offset -= 40

	// get locations
	resp, err := cfg.PokedexClient.ListLocationAreas(cfg.offset)
	if err != nil {
		return err
	}

	results := resp.Results

	// print areas
	fmt.Printf("\nThere are %d areas!\n", resp.Count)
	fmt.Printf("Page: %d. Location areas:\n", cfg.offset/20)
	fmt.Println()
	for i, result := range results {
		fmt.Printf(" %04d. %s\n", cfg.offset+i+1, result.Name)
	}
	fmt.Println()

	// config for next listing
	cfg.offset += 20
	cfg.prevCommand = "mapb"

	return nil
}