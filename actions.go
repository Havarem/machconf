package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/urfave/cli"

	"gopkg.in/yaml.v2"
)

const filePath = ".machconf"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeConfig(machine Machine) (error) {
	yml, err := yaml.Marshal(machine)
	if err != nil {
		return err 
	}

	err = ioutil.WriteFile(filePath, yml, 0755)
	return err
}

// Handles the init command of the machconf tool
func InitAction(c *cli.Context) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		machine := Machine {
			Ip: nil,
			Domain: "",
			Flags: []Flag { },
		}

		err = writeConfig(machine)
		check(err)

		os.Stdout.WriteString("Machine configuration created\n")
	} else {
		os.Stderr.WriteString("Can't create machine - Configuration already exists\n")
	}

	return nil
}

func openConfig() (Machine, error) {
	var result Machine
	var err error

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Stderr.WriteString("Machine not initialized")
	} else {
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return result, err
		}

		err = yaml.Unmarshal(content, &result)
	}

	return result, err
}

// Handles the configure command of the machconf tool
func ConfigureAction(c *cli.Context) error {
	machine, err := openConfig()
	check(err)

	if c.String("fqdn") != "" {
		machine.Domain = c.String("fqdn")
	}

	if c.String("ip") != "" {
		machine.Ip = net.ParseIP(c.String("ip"))
	}

	err = writeConfig(machine)
	check(err)

	os.Stdout.WriteString("Machine updated\n")

	return nil
}

// Handles the tag command of the machconf tool
func TagAction(c *cli.Context) error {
	fmt.Println("Tag not yet implemented")
	return nil
}