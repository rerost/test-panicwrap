package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/panicwrap"
)

func main() {
	exitStatus, err := panicwrap.BasicWrap(panicHandler)
	if err != nil {
		panic(err)
	}

	if exitStatus >= 0 {
		os.Exit(exitStatus)
	}

	if !panicwrap.Wrapped(nil) {
		// Unreachable code
		os.Exit(1)
	}

	run()
}

func run() {
	fmt.Println(os.Getpid())
	time.Sleep(1 * time.Minute)
	os.Exit(0)
}

func panicHandler(output string) {
	fmt.Fprintf(os.Stderr, output)
}
