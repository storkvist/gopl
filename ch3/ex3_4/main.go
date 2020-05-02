package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		Surface(w, r.FormValue("fill"))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
