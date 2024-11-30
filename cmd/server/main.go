package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/icchy-san/uctest/config"
	"github.com/icchy-san/uctest/db"
	"github.com/icchy-san/uctest/server"
	"github.com/icchy-san/uctest/service"
	"golang.org/x/sync/errgroup"
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

	app := server.New(ctx, invoiceService)

	// Start server with error group to track all goroutine error
	//   and stop all of them when one of them is crashed.
	eg, ectx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return app.Listen(":8080")
	})

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	select {
	case <-signalChannel:
		log.Println("received signal for killing the process")
	case <-ectx.Done():
	}

	return exitOK
}
