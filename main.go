package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	inputFilename := flag.String("if", "", "input filename")
	outputFilename := flag.String("of", "", "output filename")

	flag.Parse()

	if *inputFilename == "" || *outputFilename == "" {
		flag.PrintDefaults()
		os.Exit(2)
	}

	inputData, err := os.ReadFile(*inputFilename)
	exitf("could not read an input file", err)

	input := &Input{}

	err = yaml.Unmarshal(inputData, input)
	exitf("could not parse input data", err)

	output, err := os.Create(*outputFilename)
	exitf("could not create an output file", err)

	defer output.Close()

	header := "package " + input.Package + "\n\n"

	if len(input.Imports) > 0 {
		header += "import (\n"
		for _, imp := range input.Imports {
			header += "\t\"" + imp + "\"\n"
		}
		header += ")\n\n"
	}

	header += "// ATTENTION: Do not change this file manually. This file was generated via errgen utility.\n"

	_, err = output.WriteString(header)
	exitf("could not put a header into output file", err)

	for i := 0; i < len(input.Items); i++ {
		item := input.Items[i]
		var itemString string

		if item.Desc != "" {
			itemString = "\n// " + input.Prefix + item.Name + " " + item.Desc + ".\n"
		}

		var groupName string

		if item.Group != "" {
			groupName, err = mapGroup(input.GroupPrefix, input.Groups, item.Group)
			exitf("could not get a group name", err)
		}

		itemString += "var " + input.Prefix + item.Name + " = " + input.Struct + "{Code: " + fmt.Sprintf("%d", i+input.Start) +
			", Text: \"" + item.Text + "\""

		if groupName != "" {
			itemString += ", Group: " + groupName
		}

		itemString += "}\n"

		_, err := output.WriteString(itemString)
		exitf("cound not put an item into output file", err)
	}

	output.Sync()
}

func exitf(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+" error=[%v]\n", err)
		os.Exit(1)
	}
}

func mapGroup(prefix string, groups []Group, tag string) (string, error) {
	for i := 0; i < len(groups); i++ {
		if groups[i].Tag == tag {
			return prefix + groups[i].Name, nil
		}
	}

	return "", fmt.Errorf("unexpected group tag: %s", tag)
}
