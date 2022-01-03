package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

// HealthCheck API returns date time to client
// func HealthCheck(w http.ResponseWriter, req *http.Request) {
// 	currentTime := time.Now()
// 	io.WriteString(w, currentTime.String())
// }

type Router struct{}

func NewRouter() *Router { return &Router{} }
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {

	l := log.New(os.Stdout, " api idfs ", log.LstdFlags)
	rt := NewRouter()
	srv := http.Server{
		Addr:         "127.0.0.1:8000", // configure the bind address
		Handler:      rt,               // set the default handler
		ErrorLog:     l,                // set the logger for the server
		ReadTimeout:  0 * time.Second,  // max time to read request from the client
		WriteTimeout: 1 * time.Second,  // max time to write response to the client
		IdleTimeout:  30 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// rt.HandleFunc("/idfs", handlers.GetIdfs).Methods("GET")
	// rt.HandleFunc("/idfs/{id}", handlers.GetIdfByID).Methods("GET")
	// rt.HandleFunc("/parts", handlers.GetParts).Methods("GET")
	// rt.HandleFunc("/parts/{id}", handlers.GetPartByID).Methods("GET")
	// rt.HandleFunc("/status", handlers.HealthCheck).Methods("GET")
	log.Fatal(srv.ListenAndServe())
}
