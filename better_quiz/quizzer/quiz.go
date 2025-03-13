package quizzer

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"strings"
	"time"
)

type Data struct {
	question string
	answer   string
}

func Init() {
	f := openFile("quiz.txt")
	defer f.Close()
	ch := make(chan string, 1)
	score := 0
	lines := 0
	for d := range iterFile(f) {
		go readAnswer(d.question, ch)
		t := time.NewTimer(10 * time.Second)
		ans, isRcx := isAnswerReceived(ch, t)
		if isRcx {
			if ans == d.answer {
				score++
			}
		}
		lines++
	}
	fmt.Printf("### Your Score: %d/%d\n", score, lines)
}

func isAnswerReceived(ch <-chan string, t *time.Timer) (string, bool) {
	select {
	case ans := <-ch:
		{
			fmt.Printf("answer received from user: %s\n", ans)
			return ans, true
		}
	case ct := <-t.C:
		fmt.Printf("time out: %v\n", ct)
		return "", false
	}
}

func readAnswer(q string, ch chan string) {
	fmt.Println("#Question: ", q)
	fmt.Println("Pls provide an answer")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	ans := sc.Text()
	ch <- ans
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
			d := Data{ss[0], ss[1]}
			if !yield(d) {
				break
			}
		}
	}
}
