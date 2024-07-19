package commands

import (
	"fmt"
	"os"
)

func Exit() {
	fmt.Print("Exiting...")
	os.Exit(0)
}