package main

import "github.com/urfave/cli"

var (
	InitCmd = cli.Command {
		Name: "init",
		Usage: "Initialize a machine configuration file",
		Action: InitAction,
	}
	ConfigureCmd = cli.Command {
		Name: "configure",
		Usage: "Configure the machine",
		Flags: []cli.Flag { FqdnFlag, Ip4Flag, TagsFlag },
		Action: ConfigureAction,
	}
	TagCmd = cli.Command {
		Name: "tag",
		Usage: "Manages machine configuration tags",
		Flags: []cli.Flag { TagsFlag, RemoveTagFlag },
		Action: TagAction,
	}
)