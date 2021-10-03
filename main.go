package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/blog/hello-from-go", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello from Go!"))
	})

	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {

		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		rw.Write([]byte("Your id is " + id + " and your name is " + name))
	})

	http.HandleFunc("/logout", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Logging out!"))
	})

	log.Println("Starting servre on port 8080")
	http.ListenAndServe(":8080", nil)
}









