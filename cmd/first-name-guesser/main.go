package main

import (
	"first-name-guesser"
	"fmt"
	"log"
	"net/http"
	"os"
)

func nameHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	log.Println("Got name request!", name)
	log.Println("GUESS:", guesser.GuessFirstName(name))
	w.Write([]byte(fmt.Sprintf(`{"first_name": "%s"}`, guesser.GuessFirstName(name))))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/name", nameHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s\n", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
