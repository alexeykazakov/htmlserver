package main

import (
	"context"
	"fmt"
	"github.com/alexeykazakov/htmlserver/pkg/assets/server"
	"github.com/alexeykazakov/htmlserver/pkg/configuration"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	srv := server.New()

	err := srv.SetupRoutes()
	if err != nil {
		panic(err.Error())
	}

	routesToPrint := srv.GetRegisteredRoutes()
	fmt.Printf("Configured routes: %s\n", routesToPrint)

	// listen concurrently to allow for graceful shutdown
	go func() {
		fmt.Printf("Listening on %q...\n", "8090")
		if err := srv.HTTPServer().ListenAndServe(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	gracefulShutdown(srv.HTTPServer())
}

func gracefulShutdown(hs ...*http.Server) {
	// For a channel used for notification of just one signal value, a buffer of
	// size 1 is sufficient.
	stop := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM
	// (Ctrl+/). SIGKILL, SIGQUIT will not be caught.
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	sigReceived := <-stop
	fmt.Printf("Signal received: %+v\n", sigReceived.String())

	ctx, cancel := context.WithTimeout(context.Background(), configuration.GracefulTimeout)
	defer cancel()
	fmt.Printf("Shutdown with timeout: %s", configuration.GracefulTimeout)
	for _, s := range hs {
		if err := s.Shutdown(ctx); err != nil {
			fmt.Printf("Shutdown error: %s", err.Error())
		} else {
			fmt.Println("Server stopped.")
		}
	}
}
