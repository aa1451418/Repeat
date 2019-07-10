package main

import (
	"fmt"
	"net/http"
)

func indextwo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Web_Service...</h1>")
}

func main() {
	http.HandleFunc("/", indextwo)
	fmt.Println("Start_Service.....")
	http.ListenAndServe(":3214", nil)
}
