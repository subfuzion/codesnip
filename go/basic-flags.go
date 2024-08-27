package main

import (
	"flag"
	"fmt"
)

func main() {

	// (flag, default value, description)
	boolFlag := flag.Bool("bool", false, "bool flag")
	intFlag := flag.Int("int", 0, "int flag")
	stringFlag := flag.String("string", "", "string flag")

	flag.Parse()

	fmt.Println("boolFlag (--bool)     :", *boolFlag)
	fmt.Println("intFlag (--int)       :", *intFlag)
	fmt.Println("stringFlag (--string) :", *stringFlag)
}
