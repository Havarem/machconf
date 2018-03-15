package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Generic Machine Configuration"
	app.Version = "0.0.1-SNAPSHOT"
	app.Usage = "Handles Generic Machine Configuration system"
	app.Commands = []cli.Command { InitCmd, ConfigureCmd, TagCmd }

	app.Run(os.Args);
}