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
)

func main() {
	l := log.New(os.Stdout, " api idfs ", log.LstdFlags)
	// create a new serve mux and register the handlers
	sm := http.NewServeMux()

	// create a new server
	s := http.Server{
		Addr:         "127.0.0.1:8001",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	sm.HandleFunc("/idfs", handlers.GetIdfs)
	sm.HandleFunc("/idfs/", handlers.GetIdfByID)
	sm.HandleFunc("/parts", handlers.GetParts)
	sm.HandleFunc("/parts/", handlers.GetPartByID)
	sm.HandleFunc("/status", handlers.HealthCheck)
	sm.HandleFunc("/slow", handlers.Slow)

	l.Println("[INFO] Starting server on port 8001")
	//start the server
	go func() {
		log.Fatal(s.ListenAndServe())
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	l.Println("[INFO] Press Ctrl-C to stop service")
	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGINT)

	// Block until a signal is received.
	sig := <-c
	log.Println("[INFO] Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancelCtx := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancelCtx()
	err := s.Shutdown(ctx)
	if err != nil {
		l.Printf("[ERROR] Error starting shutdowns: %s\n", err)
		os.Exit(1)
	}
}
