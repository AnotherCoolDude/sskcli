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

// Split splits the config struct into parts
func (c *Config) Split() (tools []string, dependencies map[string][]string, desc map[string][]string) {
	toolDeps := map[string][]string{}
	t := []string{}
	descs := map[string][]string{}
	for _, tool := range c.Tools {
		t = append(t, tool.Name)
		actions := []string{}
		for _, action := range tool.Actions {
			actions = append(actions, action.Name)
			descs[action.Name] = []string{action.Description}
		}
		toolDeps[tool.Name] = actions
	}
	return t, toolDeps, descs
}
