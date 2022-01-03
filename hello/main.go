package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/elkmos/my_go/handlers"
)

// HealthCheck API returns date time to client
// func HealthCheck(w http.ResponseWriter, req *http.Request) {
// 	currentTime := time.Now()
// 	io.WriteString(w, currentTime.String())
// }

func main() {

	// config du logger
	l := log.New(os.Stdout, " api idfs ", log.LstdFlags)

	// create new handler
	myh := handlers.NewHandler(l)
	//  create new serverMux
	sm := http.NewServeMux()

	//  route to myhandler
	sm.Handle("/idfs/", myh)

	// create custum  http server  http
	server := http.Server{
		Addr:         "127.0.0.1:8000",  // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	l.Println("[INFO] Starting server on port 8000")
	log.Fatal(server.ListenAndServe())
	// trap sigterm or interupt and gracefully shutdown the server
	l.Println("[INFO] Press Ctrl-C to stop service")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("[INFO] Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
