package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	echo := "job started"
	successMsg := "job succeeded"
	failMsg := "job failed"
	exitCode := 1
	fail := false

	flag.StringVar(&echo,
		"echo", echo, "initial message")
	flag.StringVar(&successMsg,
		"success-msg", successMsg, "success message")
	flag.StringVar(&failMsg,
		"fail-msg", failMsg, "failure message")
	flag.IntVar(&exitCode,
		"exitcode", exitCode, "failure exit code (default=1)")
	flag.BoolVar(&fail,
		"fail", fail, "make process fail")

	flag.Parse()

	if echo != "" {
		fmt.Println(echo)
	}

	if fail {
		fmt.Println(failMsg)
		os.Exit(exitCode)
	}

	fmt.Println(successMsg)
}
