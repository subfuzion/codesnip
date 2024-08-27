package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Buffered input that splits input on lines
	input := bufio.NewScanner(os.Stdin)

	// Buffered output
	output := bufio.NewWriter(os.Stdout)

	lineNo := 0

	// Scan until no more input
	for input.Scan() {
		text := input.Text()
		lineNo++
		s := fmt.Sprintf("%03d %s\n", lineNo, text)

		// Would be simple and straightforward to just use fmt.Println
		// here, but want to emphasize piping from stdin to stdout
		// explicitly. Intentionally ignoring return values for demo.
		_, _ = output.WriteString(s)
		_ = output.Flush()
	}

}
