package main

import "net/http"

func main() {
	http.HandleFunc("/", FindZipCodeCEP)
	http.ListenAndServe(":8080", nil)
}

func FindZipCodeCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
