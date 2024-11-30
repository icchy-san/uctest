package main

import (
	"context"
	"fmt"
	"github.com/icchy-san/uctest/config"
	"github.com/icchy-san/uctest/db"
	"github.com/icchy-san/uctest/server"
	"github.com/icchy-san/uctest/service"
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
	env, err := config.ReadFromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading env: %v\n", err)
		return exitError
	}

	database, initDBErr := db.New(env)
	if initDBErr != nil {
		fmt.Fprintf(os.Stderr, "Error initializing database: %v\n", initDBErr)
		return exitError
	}

	invoiceService := service.New(database)

	// DI
	app := server.New(ctx, invoiceService)
	app.Listen(":8080")

	return exitOK
}
