package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

func postArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://www.idntimes.local")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("\n%s\n\n", string(requestDump))
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/article", postArticle).Methods("POST")
	log.Fatal((http.ListenAndServe(":8080", r)))
}
