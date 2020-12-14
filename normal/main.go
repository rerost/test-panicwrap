package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	run()
}

func run() {
	fmt.Println(os.Getpid())
	time.Sleep(1 * time.Minute)
	os.Exit(0)
}
