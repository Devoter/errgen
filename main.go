package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

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

	inputData, err := ioutil.ReadFile(*inputFilename)
	exitf("could not read an input file", err)

	input := &Input{}

	err = yaml.Unmarshal(inputData, input)
	exitf("could not parse input data", err)

	output, err := os.Create(*outputFilename)
	exitf("could not create an output file", err)

	defer output.Close()

	header := "// Do not change this file manually. This file was generated via errgen utility.\npackage \"" + input.Package +
		"\"\n\n"

	_, err = output.WriteString(header)
	exitf("coult not put a header into output file", err)

	for i := 0; i < len(input.Items); i++ {
		item := input.Items[i]
		itemString := "// Error" + item.Name + " " + item.Desc + ".\n" +
			"var Error" + item.Name + " = ExtenedError{Code: " + strconv.FormatInt(int64(i+input.Start), 10) + ", Text: \"" +
			item.Text + "\"}\n\n"
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
