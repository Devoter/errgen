package main

// Input is an input document.
type Input struct {
	Start   int      `yaml:"start"`
	Package string   `yaml:"package"`
	Struct  string   `yaml:"struct"`
	Prefix  string   `yaml:"prefix"`
	Imports []string `yaml:"imports"`
	Items   []*Item  `yaml:"items"`
}
