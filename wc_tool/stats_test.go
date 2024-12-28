package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestCalculateStats(t *testing.T) {
	tc := []struct {
		name  string
		input string
		want  Stats
	}{
		{
			name:  "Single Line",
			input: "Hello World",
			want: Stats{
				LineCount: 1,
				WordCount: 2,
				CharCount: 10,
				ByteCount: 10,
			},
		},
		{
			name:  "Empty",
			input: "",
			want: Stats{
				LineCount: 0,
				WordCount: 0,
				CharCount: 0,
				ByteCount: 0,
			},
		},
		{
			name:  "Space Character",
			input: "\n",
			want: Stats{
				LineCount: 1,
			},
		},
	}

	for _, rt := range tc {
		t.Run(rt.name, func(t *testing.T) {
			got := CalculateStats(bufio.NewReader(strings.NewReader(rt.input))) // prints to the console
			if got != rt.want {
				t.Errorf("got: %v, want: %v", got, rt.want)
			}
		})
	}
}
