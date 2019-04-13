package main

import (
	"fmt"
	"log"
	"net/http"
)

func statusCode(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I'm ok!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<div align='center'><img src='static/assets/cat.gif' alt='小猫咪' /></div>")
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/status", statusCode)

	http.Handle("/static/assets/", http.StripPrefix("/static/assets/", http.FileServer(http.Dir("./static/assets"))))
	fmt.Println("Listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
