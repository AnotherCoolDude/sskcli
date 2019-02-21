package parser

type Config struct {
	Title string `yaml:"toolListTitle"`
	Tools struct {
		Name    string `yaml:"name"`
		Actions struct {
			Name        string `yaml:"name"`
			Description string `yaml:"desc"`
		}
	}
}
