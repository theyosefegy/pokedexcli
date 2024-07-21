package commands

import (
	"fmt"
	"os"
)

func Exit(cfg *Config) {
	fmt.Print("Exiting...")
	os.Exit(0)
}