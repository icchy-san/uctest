package main

import (
	"context"
	"github.com/icchy-san/uctest/server"
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
	ctx := context.Background()
	//env, err := config.ReadFromEnv()
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error reading env: %v\n", err)
	//	return exitError
	//}

	app := server.New(ctx)
	app.Listen(":8080")

	return exitOK
}
