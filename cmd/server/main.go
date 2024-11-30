package main

import (
	"fmt"
	"github.com/icchy-san/uctest/config"
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
	env, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading env: %v\n", err)
		return exitError
	}

	fmt.Fprintf(os.Stdout, "%v", env)
	return exitOK
}
