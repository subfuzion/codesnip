package main

import (
	"fmt"
	"os"
)

func main() {

	// Take a slice of args starting at index 1
	// to skip the program name.
	var args = os.Args[1:]

	// Print each arg on its own line.
	for i, arg := range args {
		fmt.Println(i+1, arg)
	}
}
