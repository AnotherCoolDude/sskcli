package parser

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// ParseYaml parses the yaml file at path path
func ParseYaml(path string) *Config {
	c := &Config{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading file at %s: %s\n", path, err)
	}
	yaml.Unmarshal(data, c)
	fmt.Printf("Config: %v\n", c)
	return c
}
