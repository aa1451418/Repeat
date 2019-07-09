package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello WhatsApp Number</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Start_Serverce...")
	http.ListenAndServe(":2435", nil)
}
