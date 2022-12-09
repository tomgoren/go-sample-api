package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func validateContentType(contentTypeHeader string) (valid bool, err error) {
	if contentTypeHeader == "application/json" {
		return true, nil
	} else {
		return false, errors.New("Content-Type header missing or unsupported type passed")
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	contentTypeHeader := r.Header.Get("Content-Type")
	valid, err := validateContentType(contentTypeHeader)
	if err != nil {
		log.Print(err)
		http.Error(w, "Bad request - missing correct Content-Type header, expected 'application/json'", 400)
	}
	requestContent, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if valid {
		log.Printf("Received input: %v at URL: %v", string(requestContent), r.URL)
		w.Write([]byte(`{"message": "hello world"}`))
	}
}
