package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hey there mod users")

}

// w is an http response writer
func serveHome(w http.ResponseWriter, r *http.Request) {
	// We need to send some response to a request we receiveconst
	w.Write([]byte("<h1>Welcome to Mod chapter</h1>"))

}
