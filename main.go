package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	http.FileServer(http.Dir("./static/"))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
