package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"unicode"
)

type Options struct {
	PrintLines bool
	PrintWords bool
	PrintChars bool
	PrintBytes bool
}

type Stats struct {
	WordCount uint64
	LineCount uint64
	CharCount uint64
	ByteCount uint64
}

func CalculateStats(reader *bufio.Reader) Stats {
	stat := Stats{}
	var prevChar rune
	for {
		r, s, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				if prevChar != rune(0) && unicode.IsSpace(r) {
					stat.WordCount++
					break
				}
			}
			log.Fatalf("unexpected error encountered: %v\n", err)
		}
		if r == '\n' {
			stat.LineCount++
		}
		if !unicode.IsSpace(prevChar) && unicode.IsSpace(r) {
			stat.WordCount++
		}
		stat.CharCount++
		stat.ByteCount += uint64(s)
	}

	return stat
}

func printOutput(stats Stats, options Options) {
	if options.PrintLines {
		fmt.Printf("%d ", stats.LineCount)
	}
	if options.PrintWords {
		fmt.Printf("%d ", stats.WordCount)
	}
	if options.PrintChars {
		fmt.Printf("%d ", stats.CharCount)
	}
	if options.PrintBytes {
		fmt.Printf("%d ", stats.ByteCount)
	}
	fmt.Println()
}

// runes are int representation of UTF-8 characters.
func countChars(ss []string) int {
	accumulator := 0
	for _, word := range ss {
		r := []rune(word)
		accumulator += len(r)
	}
	return accumulator
}

// bytes are int representation of ASCII characters.
func countBytes(ss []string) int {
	accumulator := 0
	for _, word := range ss {
		r := []byte(word)
		accumulator += len(r)
	}

	return accumulator
}
