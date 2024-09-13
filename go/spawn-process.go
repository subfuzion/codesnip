package main

import (
	"fmt"
	"os/exec"
)

func run(msg string, succeed bool) {

	// This demo requires `./job-process` to exist in the cwd.
	// If not, go build it: `go build job-process.go`
	job := "./job-process"

	var cmd *exec.Cmd

	if succeed {
		cmd = exec.Command(job)
	} else {
		cmd = exec.Command(job,
			"--fail",
			"--fail-msg='RUHROH: catastrophic mission failure'",
			"--exitcode=2")
	}

	fmt.Printf("*** Spawning job: %s (should succeed: %t)\n",
		msg, succeed)

	out, err := cmd.CombinedOutput()

	// Whether successful or not, print all the output
	fmt.Printf("%s\n", out)

	// If not successful, also print the error code
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	run("good job", true)
	run("bad job", false)
}
