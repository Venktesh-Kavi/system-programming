package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const (
	LINES = "l"
	WORDS = "w"
	CHARS = "c"
	BYTES = "b"
)

func Run() {
	configOptions := Options{}
	flag.BoolVar(&configOptions.PrintLines, LINES, false, "Count lines")
	flag.BoolVar(&configOptions.PrintWords, WORDS, false, "Count Words")
	flag.BoolVar(&configOptions.PrintChars, CHARS, false, "Count Chars")
	flag.BoolVar(&configOptions.PrintBytes, BYTES, false, "Count Bytes")
	flag.Parse()

	if !configOptions.PrintLines && !configOptions.PrintWords && !configOptions.PrintChars {
		configOptions.PrintLines = true
		configOptions.PrintWords = true
		configOptions.PrintChars = true
	}

	commandLineArgs := flag.CommandLine.Args()
	handleInput(commandLineArgs, configOptions)
}

func handleInput(args []string, configOptions Options) {
	if len(args) == 0 {
		// check for stdin
		r := bufio.NewReader(os.Stdin)
		stats := CalculateStats(r)
		printOutput(stats, configOptions)
	} else {
		for _, arg := range args {
			stats := processFile(arg)
			printOutput(stats, configOptions)
		}
	}
}

func processFile(arg string) Stats {
	file, err := os.Open(arg)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("error closing file: %v\n", err)
		}
	}(file)
	return CalculateStats(bufio.NewReader(file))
}
