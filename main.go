package main

import (
	"os"

	"github.com/MerlinDMC/jsonbeat/cmd"
	_ "github.com/MerlinDMC/jsonbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
