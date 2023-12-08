package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Useme landing page!\n")
	}

	http.HandleFunc("/useme", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/useme")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
