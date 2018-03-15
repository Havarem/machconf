package main

import "github.com/urfave/cli"

var (
	// Fully Qualified Domain Name flag definition
	FqdnFlag = cli.StringFlag {
		Name: "fqdn, f",
		Usage: "Sets the `Fully Qualified Domain Name`",
	}
	// IPv4 flag definition
	Ip4Flag = cli.StringFlag {
		Name: "ip4",
		Usage: "Sets the `IPv4 address`",
	}
	// Add Tags flag definition
	TagsFlag = cli.StringFlag {
		Name: "add-tags, t",
		Usage: "Adds `tag(s)` to the machine",
	}
	// Remove tag flag definiion
	RemoveTagFlag = cli.StringFlag {
		Name: "remove-tag",
		Usage: "Removes a `tag` to the machine",
	}
)