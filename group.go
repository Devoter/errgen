package main

// Group pair: original name and value.
type Group struct {
	Name string `yaml:"name"` // original group value
	Tag  string `yaml:"tag"`  // name which is used in yaml input
}
