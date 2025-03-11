package quizzer

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"strings"
)

type Data struct {
	question string
	answer   string
}

func Init() {
	fmt.Println("reading quiz file")
	f := openFile("quiz.txt")
	defer f.Close()
	for d := range iterFile(f) {
		fmt.Println(d.question)
	}
}

func openFile(fn string) *os.File {
	f, err := os.Open(fn)

	if err != nil {
		log.Fatalf("unable to open quiz file: %v\n", err)
	}

	return f
}

// iterFile returns a generator function, which read a line from a file and produces a data object.
// The generator function encapsulates the logic for reading the file and transforming into the data object.
// The yield function is provided by the iter package, it is used by the generator function to provide data back to the iterator (the range loop). The yield fn determines whether the generator should continue or not.
func iterFile(f *os.File) iter.Seq[Data] {
	return func(yield func(Data) bool) {
		sc := bufio.NewScanner(f)
		sc.Split(bufio.ScanLines)
		for sc.Scan() {
			line := sc.Text()
			if line == "" {
				fmt.Println("empty line")
			}
			ss := strings.Split(line, ",")
			fmt.Println("read line: ", line)
			d := Data{ss[0], ss[1]}

			if !yield(d) {
				break
			}
		}
	}
}
