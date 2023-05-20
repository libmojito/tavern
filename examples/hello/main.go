package main

import (
	"os"

	"github.com/libmojito/tavern/examples/hello/cmd"
)

func main() {
	err := cmd.NewCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
