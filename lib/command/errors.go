package command

import "fmt"

type InvalidCommandOptionType struct {
	OptionType int
	Path       string
}

type UnknownCommandOptionType struct {
	InvalidType int
	Path        string
}

type MissingCommandField struct {
	FieldName string
	Path      string
}

type InvalidCommandOptionChoiceValue struct {
	Expected string
	Actual   string
	Path     string
}

type CommandOptionFieldNotAllowed struct {
	Field string
	Type  int
	Path  string
}

type CommandFieldExceedsMaxLen struct {
	Field string
	Found int
	Max   int
	Path  string
}

func (e InvalidCommandOptionType) Error() string {
	return fmt.Sprintf("invalid command option `%v` at `%v`", e.OptionType, e.Path)
}

func (e UnknownCommandOptionType) Error() string {
	return fmt.Sprintf("unknown command option type `%v` at `%v`", int(e.InvalidType), string(e.Path))
}

func (e MissingCommandField) Error() string {
	return fmt.Sprintf("missing required field `%v` at `%v`", string(e.FieldName), string(e.Path))
}

func (e CommandOptionFieldNotAllowed) Error() string {
	return fmt.Sprintf("option type `%v` cannot have field `%v` at `%v`", e.Type, e.Field, e.Path)
}

func (e InvalidCommandOptionChoiceValue) Error() string {
	return fmt.Sprintf("expected choice.value type `%v`, got `%v` at `%v`", e.Expected, e.Actual, e.Path)
}

func (e CommandFieldExceedsMaxLen) Error() string {
	return fmt.Sprintf("field `%v` exceeds maximum length `%v`, found `%v` at `%v`", e.Field, e.Max, e.Found, e.Path)
}
