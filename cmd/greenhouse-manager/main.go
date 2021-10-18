package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KaiserWerk/Greenhouse-Manager/internal/handler"

	"github.com/gorilla/mux"
)

var (
	port int
)

func main() {
	flag.IntVar(&port, "port", 47336, "The port to listen on")
	flag.Parse()

	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("receive", handler.ReceiveHandler).Methods(http.MethodPost)

	router.HandleFunc("/", handler.IndexHandler)

	srv := http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%d", port),
		ReadTimeout: 20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt)

	go func() {
		<-exitCh
		fmt.Println("server shutdown started")

		// do other stuff

		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			fmt.Printf("could not shut down server: %s\n", err.Error())
			os.Exit(-1)
		}
	}()

	fmt.Printf("server listening on :%d...\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server: %s\n", err.Error())
		os.Exit(-2)
	}

	fmt.Println("server shutdown complete")
}


