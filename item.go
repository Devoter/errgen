package main

// Item is an error item.
type Item struct {
	Name  string `yaml:"name"`
	Code  int    `yaml:"-"`
	Text  string `yaml:"text"`
	Group string `yaml:"group,omitempty"`
	Desc  string `yaml:"desc,omitempty"`
}
