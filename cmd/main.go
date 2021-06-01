package main

import (
	"log"
	"net/http"
)

func main()  {
	log.Println("starting server...")
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/JSON")
		w.Write([]byte(`{"response": "success"}`))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
