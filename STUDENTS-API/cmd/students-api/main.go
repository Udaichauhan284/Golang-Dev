package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Udaichauhan284/Golang-Dev/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Students APIs"))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started: ", slog.String("address", cfg.Addr));

	// fmt.Printf("Server started %s", cfg.HTTPServer.Addr)

	//Now implementing the gracefully interupt, because if we not doing that in production, when there is ongoing task, it will stop badly.
	//to implement the gracefully stop, use the go routine.

	//making the channel
	done := make(chan os.Signal, 1);

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM);

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server");

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second);
	defer cancel()

	err := server.Shutdown(ctx);
	if err != nil {
		slog.Error("Failed to Shutdown server ", slog.String("error", err.Error()));
	}

	slog.Info("Server Shutdown successfully");
}
