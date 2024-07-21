package commands

import (
	"fmt"
	"os"
)

func Exit(cfg *Config) error{
	fmt.Print("Exiting...")
	os.Exit(0)
	return nil
}