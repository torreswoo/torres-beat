package main

import (
	"os"

	"github.com/torreswoo/torresbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
