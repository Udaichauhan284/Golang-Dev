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
	student "github.com/Udaichauhan284/Golang-Dev/internal/http/handlers/students"
	"github.com/Udaichauhan284/Golang-Dev/internal/storage/sqlite"
)

func main() {
	//load config
	cfg := config.MustLoad()

	//database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage));

	//now creating the router to getById
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage));

	//now creating the endpoint to access the list of student
	router.HandleFunc("GET /api/students", student.GetList(storage));

	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudent(storage));

	router.HandleFunc("DELETE /api/students/{id}", student.DeleteStudent(storage));

	//setup server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started: ", slog.String("address", cfg.Addr))

	// fmt.Printf("Server started %s", cfg.HTTPServer.Addr)

	//Now implementing the gracefully interupt, because if we not doing that in production, when there is ongoing task, it will stop badly.
	//to implement the gracefully stop, use the go routine.

	//making the channel
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err1 := server.Shutdown(ctx)
	if err1 != nil {
		slog.Error("Failed to Shutdown server ", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown successfully")
}
