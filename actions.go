package main

import (
	"io/ioutil"
	"net"
	"os"
	"strings"

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
			Tags: []Tag { },
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

func unique(tags []Tag) []Tag {
	u := make([]Tag, 0, len(tags))
	m := make(map[Tag]bool)

	for _, val := range tags {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func separateTags(tags string) []Tag {
	var result []Tag
	resultString := strings.Split(tags, ",")

	for i := 0; i < len(resultString); i++ {
		result = append(result, Tag(resultString[i]))
	}

	return result
}

// Handles the configure command of the machconf tool
func ConfigureAction(c *cli.Context) error {
	machine, err := openConfig()
	check(err)

	if c.String("fqdn") != "" {
		machine.Domain = c.String("fqdn")
	}

	if c.String("ip4") != "" {
		machine.Ip = net.ParseIP(c.String("ip4"))
	}

	if c.String("add-tags") != "" {
		addTag := separateTags(c.String("add-tags"))

		for i := 0; i < len(addTag); i++ {
			machine.Tags = unique(append(machine.Tags, addTag[i]))
		}
	}

	err = writeConfig(machine)
	check(err)

	os.Stdout.WriteString("Machine updated\n")

	return nil
}

func difference(current []Tag, removed []Tag) []Tag {
	diffStr := []Tag {}
	m := map [Tag] int{}

	for _, s1Val := range current {
		m[s1Val] = 1
	}

	for _, s2Val := range removed {
		m[s2Val] = m[s2Val] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}

// Handles the tag command of the machconf tool
func TagAction(c *cli.Context) error {
	machine, err := openConfig()
	check(err)

	if c.String("add-tags") != "" {
		addTag := separateTags(c.String("add-tags"))

		for i := 0; i < len(addTag); i++ {
			machine.Tags = unique(append(machine.Tags, addTag[i]))
		}
	}

	if c.String("remove-tag") != "" {
		machine.Tags = difference(machine.Tags, separateTags(c.String("remove-tag")))
	}

	return nil
}