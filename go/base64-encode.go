package main

import (
	"bufio"
	"encoding/base64"
	"os"
)

func main() {

	// Buffered input that splits input on lines.
	input := bufio.NewScanner(os.Stdin)

	// Base64 Encoder/writer.
	encoder := base64.NewEncoder(
		base64.StdEncoding,
		os.Stdout)

	// Scan until EOF (no more input).
	for input.Scan() {
		bytes := input.Bytes()
		_, _ = encoder.Write(bytes)
		_, _ = encoder.Write([]byte{'\n'})
	}

	// Close the encoder and ensure it flushes remaining output
	_ = encoder.Close()
}
