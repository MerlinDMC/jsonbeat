package cmd

import (
	"github.com/MerlinDMC/jsonbeat/beater"

	"github.com/elastic/beats/libbeat/cmd"
	"github.com/elastic/beats/libbeat/cmd/instance"
)

// Name of this beat
var Name = "jsonbeat"
var Version = ""

// RootCmd to handle beats cli
var RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{Name: Name, Version: Version})
