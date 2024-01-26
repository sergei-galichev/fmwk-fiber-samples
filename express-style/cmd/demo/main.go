package main

import (
	"context"
	"express-style/internal/app"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"
)

var (
	signals = []os.Signal{
		syscall.SIGABRT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		os.Interrupt,
		syscall.SIGTERM,
	}
)

func main() {
	ctx := context.Background()
	a := app.NewApp(ctx)
	a.RunApp()

	shutdown(signals)
}

func shutdown(signals []os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	sig := <-ch
	log.Infof("Caught signal: %s. Shutting down...", sig)
}
