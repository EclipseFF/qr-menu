package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	app := NewApp(":4002", "dbdsn")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := app.Server.Start(app.Cfg.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Server.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := app.Server.Shutdown(ctx); err != nil {
		app.Server.Logger.Fatal(err)
	}
}
