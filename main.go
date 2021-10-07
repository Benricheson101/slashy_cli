package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benricheson101/slashy_cli/lib/command"
	"github.com/goccy/go-yaml"
)

func main() {
  var cmds command.CommandFile

	_cmds, err := os.ReadFile("./cmd.yml")
	if err != nil {
		log.Fatal(err)
	}

	if err := yaml.Unmarshal([]byte(_cmds), &cmds); err != nil {
		log.Fatal(err)
	}

	for _, cmd := range cmds.Commands {
		errs := cmd.Validate()

		if len(errs) != 0 {
			fmt.Println("Errors:")
			for _, e := range errs {
				fmt.Println("  =>", e)
			}

      os.Exit(1)
		}
	}
}
