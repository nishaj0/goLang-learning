package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to modules")

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	// http.ListenAndServe(":3000", r) // this may throw some err, we can use comma ok syntax, but here people mostly using log pkg
	log.Fatal(http.ListenAndServe(":3000", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang</h1>"))
}
