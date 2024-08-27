package main

import (
	"flag"
	"fmt"
)

func main() {

	boolFlag := flag.Bool("bool", false, "bool flag")
	intFlag := flag.Int("int", 0, "int flag")
	stringFlag := flag.String("string", "", "string flag")

	flag.Parse()

	fmt.Println("bool:", *boolFlag)
	fmt.Println("int:", *intFlag)
	fmt.Printf("string: \"%s\"\n", *stringFlag)
}
