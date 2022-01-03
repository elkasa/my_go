package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/elkmos/my_go/handlers"
	"github.com/gorilla/mux"
)

// HealthCheck API returns date time to client
// func HealthCheck(w http.ResponseWriter, req *http.Request) {
// 	currentTime := time.Now()
// 	io.WriteString(w, currentTime.String())
// }

func main2() {

	l := log.New(os.Stdout, " api idfs ", log.LstdFlags)
	r := mux.NewRouter()
	srv := http.Server{
		Addr:         "127.0.0.1:8000", // configure the bind address
		Handler:      r,                // set the default handler
		ErrorLog:     l,                // set the logger for the server
		ReadTimeout:  0 * time.Second,  // max time to read request from the client
		WriteTimeout: 1 * time.Second,  // max time to write response to the client
		IdleTimeout:  30 * time.Second, // max time for connections using TCP Keep-Alive
	}
	r.HandleFunc("/idfs", handlers.GetIdfs).Methods("GET")
	r.HandleFunc("/idfs/{id}", handlers.GetIdfByID).Methods("GET")
	r.HandleFunc("/parts", handlers.GetParts).Methods("GET")
	r.HandleFunc("/parts/{id}", handlers.GetPartByID).Methods("GET")
	r.HandleFunc("/status", handlers.HealthCheck).Methods("GET")
	log.Fatal(srv.ListenAndServe())
}
