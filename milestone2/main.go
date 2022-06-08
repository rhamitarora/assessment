package main

import (
	"fmt"
	"log"
	"net/http"
	t "milestone2/swt"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}	
	fmt.Fprintf(w, t.CurrentTime())
}

var serverPORT string

func main() {
	http.HandleFunc("/time", timeHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
