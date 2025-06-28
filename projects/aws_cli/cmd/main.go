package main

import (
	"acli/internal"
	"fmt"
)

func main() {
	fmt.Println("application started")
	cmd := internal.Start()
	err := cmd.Execute()
	if err != nil {
		return
	}
}
