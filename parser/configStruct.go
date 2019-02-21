package parser

// Config holds the title and the tools of config.yaml
type Config struct {
	Title string `yaml:"toolListTitle"`
	Tools []Tool `yaml:"tools"`
}

// Tool holds the title and the actions of a config tool
type Tool struct {
	Name    string   `yaml:"name"`
	Actions []Action `yaml:"actions"`
}

// Action holds the name and the description of a tool action
type Action struct {
	Name        string `yaml:"name"`
	Description string `yaml:"desc"`
}
