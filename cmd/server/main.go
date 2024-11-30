package main

import (
	"fmt"
	"os"
)

const (
	exitOK int = iota
	exitError
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(arg []string) int {
	fmt.Fprintf(os.Stdout, "called")
	return exitOK
}
