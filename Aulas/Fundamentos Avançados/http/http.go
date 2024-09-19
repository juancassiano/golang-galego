package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", chamarHome)

	http.HandleFunc("/users", chamarUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func chamarHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo"))
}

func chamarUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá usuários"))
}
