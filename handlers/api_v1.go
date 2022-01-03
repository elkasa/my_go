package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	idfs "github.com/elkmos/my_go/controllers"
	"github.com/elkmos/my_go/db"
	"github.com/gorilla/mux"
)

//  logger handler for getting and updating data
type Products struct {
	l *log.Logger
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

var errCustum = fmt.Errorf("test err: url should be /idfs/[idf]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.

func GetIdfByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idf, err := idfs.GetIdfID(id)
	if err != nil {
		fmt.Println("[ERROR] serializing idf no found", err)
		w.WriteHeader(http.StatusInternalServerError)
		db.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	db.ToJSON(idf, w)
}

func ListIdfs(w http.ResponseWriter, r *http.Request) {

	f := idfs.GetIdfs()
	w.Header().Set("Content-Type", "application/json")
	err := db.ToJSON(f, w)
	if err != nil {
		fmt.Println("[ERROR] serializing listIdf", err)
		w.WriteHeader(http.StatusInternalServerError)
		db.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	fmt.Fprintf(w, "listidf"+r.URL.Path+mux.Vars(r)["id"]+errCustum.Error())

}

func CreateIdf(w http.ResponseWriter, r *http.Request) {
	var newIdf idfs.Idf
	json.NewDecoder(r.Body).Decode(&newIdf)
	err := idfs.AddIdf(newIdf)
	if err != nil {
		fmt.Println("[ERROR] idf  creation", err)
		w.WriteHeader(http.StatusInternalServerError)
		db.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	db.ToJSON(newIdf, w)

}

func UpdateIdf(w http.ResponseWriter, r *http.Request) {
	var newIdf idfs.Idf
	json.NewDecoder(r.Body).Decode(&newIdf)
	idf, err := idfs.UpdateIdf(newIdf)
	if err != nil {
		fmt.Println("[ERROR] serializing idf no found", err)
		w.WriteHeader(http.StatusInternalServerError)
		db.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	db.ToJSON(idf, w)

}

func DeleteIdf(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idf, err := idfs.DeleteIdf(id)
	if err != nil {
		fmt.Println("[ERROR] deleting idf no found", err)
		w.WriteHeader(http.StatusInternalServerError)
		db.ToJSON(&GenericError{Message: err.Error()}, w)
	}
	db.ToJSON(idf, w)
}

// parts

func GetParts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getParts"+r.URL.Path)
}

func GetPartByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "getPartById"+r.URL.Path)

}

func Slow(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("io.ReadAll", err)

		return
	}
	defer r.Body.Close()

	fmt.Println("written")

	fmt.Fprint(w, "Hello ", string(body))
	time.Sleep(3 * time.Second)
	fmt.Fprintf(w, "H.e.l.l.o %s", string(body))
	fmt.Fprintf(w, "getPartById"+r.URL.Path)

}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "healthStatus : "+currentTime.String()+strconv.FormatInt(currentTime.Unix(), 10))

}
