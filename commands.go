package main

import "github.com/urfave/cli"

var (
	// Init command definition
	InitCmd = cli.Command {
		Name: "init",
		Usage: "Initialize a machine configuration file",
		Action: InitAction,
	}
	// Configure command definition
	ConfigureCmd = cli.Command {
		Name: "configure",
		Usage: "Configure the machine",
		Flags: []cli.Flag { FqdnFlag, Ip4Flag, TagsFlag },
		Action: ConfigureAction,
	}
	// Tag command definition
	TagCmd = cli.Command {
		Name: "tag",
		Usage: "Manages machine configuration tags",
		Flags: []cli.Flag { TagsFlag, RemoveTagFlag },
		Action: TagAction,
	}
)