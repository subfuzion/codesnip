package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Error: missing name")
		os.Exit(1)
	}

	fmt.Println("Hello", os.Args[1])
}
