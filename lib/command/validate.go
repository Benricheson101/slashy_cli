package command

import "strconv"

// TODO: read discord dev docs and finish validation
// Check if the JSON payload for creating a command is valid
func (cmd Command) Validate() []error {
	errors := make([]error, 0)

	if cmd.Name == "" {
		errors = append(errors, MissingCommandField{FieldName: "name", Path: cmd.Name + ".name"})
	} else if l := len(cmd.Name); l > 32 {
		errors = append(errors, CommandFieldExceedsMaxLen{Field: "name", Max: 32, Found: l, Path: cmd.Name + ".name"})
	}

	if cmd.Description == "" {
		errors = append(errors, MissingCommandField{FieldName: "description", Path: cmd.Name + ".description"})
	} else if l := len(cmd.Description); l > 100 {
		errors = append(errors, CommandFieldExceedsMaxLen{Field: "description", Max: 100, Found: l, Path: cmd.Name + ".description"})
	}

	validateOptions(cmd.Options, cmd.Name, &errors)

	return errors
}

// Tests if a command option is valid
func validateOptions(options []CommandOption, path string, errors *[]error) {
	for _, op := range options {
		newPath := path + "." + op.Name

	if op.Name == "" {
		*errors = append(*errors, MissingCommandField{FieldName: "name", Path: op.Name + ".name"})
	} else if l := len(op.Name); l > 32 {
		*errors = append(*errors, CommandFieldExceedsMaxLen{Field: "name", Max: 32, Found: l, Path: op.Name + ".name"})
	}

	// TODO: put this repeated code into its own function
	if op.Description == "" {
		*errors = append(*errors, MissingCommandField{FieldName: "description", Path: op.Name + ".description"})
	} else if l := len(op.Description); l > 100 {
		*errors = append(*errors, CommandFieldExceedsMaxLen{Field: "description", Max: 100, Found: l, Path: op.Name + ".description"})
	}

		if len(op.Choices) > 0 {
			if ch := len(op.Choices); ch > 25 {
				*errors = append(*errors, CommandFieldExceedsMaxLen{Field: "choices", Max: 25, Found: ch, Path: newPath})
			}

			op.validateChoices(newPath, errors)
		}

		if len(op.Options) > 0 {
			validateOptions(op.Options, newPath, errors)
		}
	}
}

// Tests if a command option choice is valid
func (c CommandOption) validateChoices(path string, errors *[]error) {
		path = path + "." + c.Name

	if c.Choices == nil {
		return
	}

	if c.Type != 3 && c.Type != 4 && c.Type != 10 {
		*errors = append(*errors, CommandOptionFieldNotAllowed{Type: c.Type, Field: "choices", Path: path})
	}

	for _, choice := range c.Choices {
    p := path + "." + choice.Name

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
