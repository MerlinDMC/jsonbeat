package main

import (
	"os"

	"github.com/MerlinDMC/jsonbeat/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
