package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello World!")
}

var serverPORT string

func main() {
	http.HandleFunc("/", helloHandler)

	if len(os.Args[1:]) < 2 {
		fmt.Printf("Missing required arguments")
		return
	} else {
		if os.Args[1] == "-port" {
			serverPORT = os.Args[2]
		}
	}

	if _, err := strconv.Atoi(serverPORT); err == nil {
		fmt.Printf("Starting server at port %s\n", serverPORT)
	} else {
		fmt.Printf("Unable to start server at port %s\n", serverPORT)
	}

	if err := http.ListenAndServe(":"+serverPORT, nil); err != nil {
		log.Fatal(err)
	}
}
