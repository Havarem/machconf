package main

import (
	"github.com/urfave/cli"
	"fmt"
	"io/ioutil"
	"os"
	"gopkg.in/yaml.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Handles the init command of the machconf tool
func InitAction(c *cli.Context) error {
	if _, err := os.Stat(".machconf"); os.IsNotExist(err) {
		machine := Machine {
			Ip: nil,
			Domain: "",
			Flags: []Flag { },
		}

		yml, err := yaml.Marshal(machine)
		check(err)

		err = ioutil.WriteFile(".machconf", yml, 0755)
		check(err)

		os.Stdout.WriteString("Machine configuration created\n")
	} else {
		os.Stderr.WriteString("Can't create machine - Configuration already exists\n")
	}

	return nil
}

// Handles the configure command of the machconf tool
func ConfigureAction(c *cli.Context) error {
	fmt.Println("Configure not yet implemented")
	return nil
}

// Handles the tag command of the machconf tool
func TagAction(c *cli.Context) error {
	fmt.Println("Tag not yet implemented")
	return nil
}