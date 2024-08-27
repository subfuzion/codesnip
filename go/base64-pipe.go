package main

import (
	"bufio"
	"encoding/base64"
	"os"
)

func main() {

	// Buffered input that splits input on lines
	input := bufio.NewScanner(os.Stdin)

	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)

	newLine := []byte("\n")

	// Scan until no more input
	for input.Scan() {
		bytes := input.Bytes()
		_, _ = encoder.Write(bytes)
		_, _ = encoder.Write(newLine)
	}
}
