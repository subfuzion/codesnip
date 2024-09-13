package main

import (
	"fmt"
	"os"
)

func main() {

	// Take a slice of args starting at index 1
	// to skip the program name.
	args := os.Args[1:]

	// Print each arg on its own line.
	for i, arg := range args {
		fmt.Printf("%d: %s\n", i+1, arg)
	}
}
