package main

// Input is an input document.
type Input struct {
	Start   int     `yaml:"start"`
	Package string  `yaml:"package"`
	Items   []*Item `yaml:"items"`
}
