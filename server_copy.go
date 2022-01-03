package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elkmos/my_go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, " api idfs ", log.LstdFlags)
	// create a new serve router and register the handlers
	router := mux.NewRouter()

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:8001",  // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	router.HandleFunc("/idfs/{id}", handlers.GetIdfByID).Methods("GET")
	router.HandleFunc("/idfs", handlers.ListIdfs).Methods("GET")
	router.HandleFunc("/idfs", handlers.CreateIdf).Methods("POST")
	router.HandleFunc("/idfs/{id}", handlers.UpdateIdf).Methods("PUT")
	router.HandleFunc("/idfs/{id}", handlers.DeleteIdf).Methods("DELETE")

	router.HandleFunc("/parts", handlers.GetParts)
	router.HandleFunc("/parts/{id}", handlers.GetPartByID)

	router.HandleFunc("/status", handlers.HealthCheck)
	router.HandleFunc("/slow", handlers.Slow)

	l.Println("[INFO] Starting server on port 8001")
	log.Fatal(s.ListenAndServe())

	// start the server

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)
	l.Println("[INFO] Press Ctrl-C to stop service")
	// Block until a signal is received.

	go func() {
		sig := <-c
		log.Println("[INFO] Got signal:", sig)
		// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
		ctx, cancelCtx := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelCtx()
		err := s.Shutdown(ctx)
		if err != nil {
			l.Printf("[ERROR] Errorshutdown server: %s\n", err)
			os.Exit(1)
		}
	}()
}
