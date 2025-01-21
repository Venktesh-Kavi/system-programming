package quizzer

import (
	"bufio"
	"fmt"
	"io"
	"iter"
	"os"
)

const path = "problems.csv"

// Init
func Init() {
	wb := []byte{}
	for line := range readQuestion(path) {
		f := os.Stdout
		wb = []byte(line)
		f.Write(wb)
		s, _ := readInput()
		fmt.Println("Read Ans: ", s)
	}
}

func readInput() (string, error) {
	f := os.Stdin
	br := bufio.NewReader(f)
	b, err := br.ReadBytes('\n')
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// read line by line with a state on till where I have read.
// readQuestion Iterator will silently skip errors
func readQuestion(path string) iter.Seq[string] {
	f, _ := os.Open(path)
	defer f.Close()
	br := bufio.NewReader(f)
	return func(yield func(string) bool) {
		for {
			bytes, err := br.ReadBytes('\n')
			if !yield(string(bytes)) {
				break
			}
			if err == io.EOF {
				break
			}
		}
	}
}
