package commands

import (
	"fmt"
	"os"
)

func Exit(cfg *Config, args ...string) error{
	fmt.Print("Exiting...")
	os.Exit(0)
	return nil
}