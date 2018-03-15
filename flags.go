package main

import "github.com/urfave/cli"

var (
	FqdnFlag = cli.StringFlag {
		Name: "fqdn, f",
		Usage: "Sets the `Fully Qualified Domain Name`",
	}
	Ip4Flag = cli.StringFlag {
		Name: "ip4",
		Usage: "Sets the `IPv4 address`",
	}
	TagsFlag = cli.StringFlag {
		Name: "add-tags, t",
		Usage: "Adds `tag(s)` to the machine",
	}
	RemoveTagFlag = cli.StringFlag {
		Name: "remove-tag",
		Usage: "Removes a `tag` to the machine",
	}
)