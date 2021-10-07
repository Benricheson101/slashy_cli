package command

import (
	"fmt"
	"strconv"
)

// Check if the JSON payload for creating a command is valid
func (cmd Command) Validate() (bool, []error) {
	errors := make([]error, 0)

	if cmd.Name == "" {
		errors = append(errors, MissingCommandField{FieldName: "name", Path: fmt.Sprintf("%v.Name", cmd.Name)})
	}

	if cmd.Description == "" {
		errors = append(errors, MissingCommandField{FieldName: "description", Path: fmt.Sprintf("%v.Description", cmd.Name)})
	}

	validateOptions(cmd.Options, cmd.Name, &errors)

	return true, errors
}

// Tests if a command option is valid
func validateOptions(options []CommandOption, path string, errors *[]error) {
	for _, op := range options {
		newPath := fmt.Sprintf("%v.%v", path, op.Name)

		op.validateChoices(newPath, errors)

		if len(op.Options) > 0 {
			validateOptions(op.Options, newPath, errors)
		}
	}
}

// Tests if a command option choice is valid
func (c CommandOption) validateChoices(path string, errors *[]error) {
	path = fmt.Sprintf("%v.%v", path, c.Name)

	if c.Choices == nil {
		return
	}

	if c.Type != 3 && c.Type != 4 && c.Type != 10 {
		*errors = append(*errors, CommandOptionFieldNotAllowed{Type: c.Type, Field: "choices", Path: path})
	}

	for _, choice := range c.Choices {
		p := fmt.Sprintf("%v.%v", path, choice.Name)

		if c.Type == 4 {
			_, err := strconv.ParseInt(choice.Value, 10, 0)

			if err != nil {
				*errors = append(*errors, InvalidCommandOptionChoiceValue{Expected: "int", Actual: "string", Path: p})
			}
		} else if c.Type == 10 {
			_, err := strconv.ParseFloat(choice.Value, 0)

			if err != nil {
				*errors = append(*errors, InvalidCommandOptionChoiceValue{Expected: "float", Actual: "string", Path: p})
			}
		}
	}
}
