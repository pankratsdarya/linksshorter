package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pankratsdarya/linksshorter/server"
	"go.uber.org/zap"
)

// сервер, на котором все крутится.

// здесь должно быть:
// запуск сервера, graceful shutdown.

func main() {
	logger, _ := zap.NewProduction()

	//nolint:errcheck Не может быть ошибки, т.к. работаем с stdout
	defer logger.Sync()

	slogger := logger.Sugar()
	slogger.Info("Starting the app...")
	slogger.Info("Reading configuration and initializing server...")

	srvr, err := server.New(slogger)
	if err != nil {
		slogger.Panic("Can't initialize server.", "err", err)
	}

	slogger.Info("Starting the server...")

	srvr.Start()

	slogger.Info("The app is ready to serve requests.")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		slogger.Infow("Received a signal.", "signal", x.String())
		// to do добавить канал с ошибками в логике
	}

	slogger.Info("Stopping the server...")

	srvr.Stop()

	slogger.Info("The app is calling the last defers and will be stopped.")
}
