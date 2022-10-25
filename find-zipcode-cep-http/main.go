package main

import "net/http"

func main() {
	http.HandleFunc("/", FindZipCodeCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func FindZipCodeCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world"))
}
