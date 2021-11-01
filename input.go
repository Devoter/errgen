package main

// Input is an input document.
type Input struct {
	Start       int      `yaml:"start,omitempty"`
	Package     string   `yaml:"package"`
	Struct      string   `yaml:"struct"`
	Prefix      string   `yaml:"prefix,omitempty"`
	Imports     []string `yaml:"imports,omitempty"`
	GroupPrefix string   `yaml:"groupPrefix,omitempty"`
	Groups      []Group  `yaml:"groups,omitempty"`
	Items       []*Item  `yaml:"items,omitempty"`
}
