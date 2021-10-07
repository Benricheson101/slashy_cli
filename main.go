package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"

	"github.com/benricheson101/slashy_cli/command"
)

type TopLevel struct {
	Commands []command.Command `yaml:"commands" json:"commands"`
}

func main() {
	var data TopLevel

	cmds, err := os.ReadFile("./cmd.yml")

	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(cmds), &data); err != nil {
		log.Fatal(err)
	}

	for _, cmd := range data.Commands {
		_, errs := cmd.Validate()

		if len(errs) != 0 {
			fmt.Println("Errors:")
			for _, e := range errs {
				fmt.Println("  =>", e)
			}
		}
	}
}
