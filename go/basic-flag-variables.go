package main

import (
	"flag"
	"fmt"
)

func main() {

	boolValue := false
	intValue := 0
	stringValue := ""

	// (flag, default value, description)
	flag.BoolVar(&boolValue, "bool", false, "bool flag")
	flag.IntVar(&intValue, "int", 0, "int flag")
	flag.StringVar(&stringValue, "string", "", "string flag")

	flag.Parse()

	fmt.Println("boolValue (--bool)     :", boolValue)
	fmt.Println("intValue (--int)       :", intValue)
	fmt.Println("stringValue (--string) :", stringValue)
}
