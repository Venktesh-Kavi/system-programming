package cli

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestCommands(t *testing.T) {
	bs := bytes.NewBufferString("")
	rootCmd.SetOut(bs)
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("test failed as root command did not execute")
	}

	bb, err := io.ReadAll(bs)
	if err != nil {
		t.Fatalf("unable to read data from the stdout of cli")
	}
	fmt.Println(string(bb))
	if len(string(bb)) == 0 {
		t.Fatalf("command line should output something")
	}
}
