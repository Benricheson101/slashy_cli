package command

type CommandFile struct {
	Commands []Command `yaml:"commands" json:"commands"`
}

type Command struct {
	Name        string          `yaml:"name" json:"name"`
	Description string          `yaml:"description" json:"description"`
	Options     []CommandOption `yaml:"options,omitempty" json:"options,omitempty"`
}

type CommandOption struct {
	Name        string                `yaml:"name" json:"name"`
	Description string                `yaml:"description" json:"description"`
	Type        int                   `yaml:"type" json:"type"`
	Required    bool                  `yaml:"required" json:"required"`
	Options     []CommandOption       `yaml:"options,omitempty" json:"options,omitempty"`
	Choices     []CommandOptionChoice `yaml:"choices,omitempty" json:"choices,omitempty"`
}

type CommandOptionChoice struct {
	Name  string `yaml:"name" json:"name"`
	Value string `yaml:"value" json:"value"`
}
