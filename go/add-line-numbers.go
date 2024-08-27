package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Buffered input that splits input on lines.
	input := bufio.NewScanner(os.Stdin)

	// Buffered output.
	output := bufio.NewWriter(os.Stdout)

	lineNo := 0

	// Scan until EOF (no more input).
	for input.Scan() {
		text := input.Text()
		lineNo++
		s := fmt.Sprintf("%03d %s\n", lineNo, text)

		// It would be simpler to just use fmt.Println,
		// but want to emphasize piping stdin to stdout
		// explicitly.
		// Intentionally ignoring return values.
		_, _ = output.WriteString(s)

	}

	// Always explicitly flush remaining buffered output.
	_ = output.Flush()
}
