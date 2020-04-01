package main

// Item is an error item.
type Item struct {
	Name string `yaml:"name"`
	Code int    `yaml:"-"`
	Text string `yaml:"text"`
	Desc string `yaml:"desc"`
}
