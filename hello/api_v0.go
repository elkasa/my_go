package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/elkmos/my_go/db"
	"github.com/gorilla/mux"
)

// idfs
type Idf struct {
	ID   string `json:"id"`
	Body string `json:"body"`
}

var ListIdf = []Idf{
	{
		ID:   "ABC00EFG",
		Body: "fname=test",
	},

	{
		ID:   "xyz00jkl",
		Body: "fname=test",
	},
}

func getIdfId(r *http.Request) int {
	id := mux.Vars(r)["id"]
	for idx := range ListIdf {
		if ListIdf[idx].ID == id {
			return idx
		}
	}
	return -1
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.

func ListIdfs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := db.ToJSON(ListIdf, w)
	if err != nil {
		fmt.Println("[ERROR] serializing listIdf", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "listidf"+r.URL.Path+mux.Vars(r)["id"])

}
func GetIdfByID(w http.ResponseWriter, r *http.Request) {
	id := getIdfId(r)
	w.Header().Set("Content-Type", "application/json")
	if id != -1 {
		err := db.ToJSON(ListIdf[id], w)
		if err != nil {
			fmt.Println("[ERROR] serializing idf no found", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		fmt.Println("[ERROR] serializing listIdf")
		http.NotFound(w, r)
		return
	}
}

func CreateIdf(w http.ResponseWriter, r *http.Request) {
	var newIdf Idf
	json.NewDecoder(r.Body).Decode(&newIdf)
	ListIdf = append(ListIdf, newIdf)
	err := db.ToJSON(ListIdf, w)
	if err != nil {
		fmt.Println("[ERROR] serializing listIdf", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func UpdateIdf(w http.ResponseWriter, r *http.Request) {
	var newIdf Idf
	id := getIdfId(r)
	fmt.Println("[info] update Idf", id)
	if id != -1 {
		json.NewDecoder(r.Body).Decode(&newIdf)
		ListIdf[id].Body = newIdf.Body
		db.ToJSON(ListIdf, w)
	} else {
		http.NotFound(w, r)
		return
	}
}
func DeleteIdf(w http.ResponseWriter, r *http.Request) {
	id := getIdfId(r)
	if id != -1 {
		ListIdf = append(ListIdf[:id], ListIdf[id+1:]...)
		db.ToJSON(ListIdf, w)
	} else {
		http.NotFound(w, r)
		return
	}
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
