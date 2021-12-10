package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/handler"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/middleware"
	"github.com/KaiserWerk/Greenhouse-Manager/internal/storage"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 47336, "The port to listen on")
	flag.Parse()

	defer func() {
		if err := storage.Close(); err != nil {
			fmt.Println("close error:", err.Error())
		}
	}()

	h := handler.HttpHandler{}
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Use(middleware.Auth)
	apiRouter.HandleFunc("/receive", h.ReceiveHandler).Methods(http.MethodPost)

	router.HandleFunc("/", h.IndexHandler)

	srv := http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	//srv.SetKeepAlivesEnabled(false)

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt)

	go func() {
		<-exitCh
		fmt.Println("server shutdown initiated")

		// do other stuff

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("could not shut down server: %s\n", err.Error())
			return
		}
	}()

	fmt.Printf("server listening on :%d...\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server: %s\n", err.Error())
		return
	}

	fmt.Println("server shutdown complete")
}
