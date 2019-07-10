package main

import "fmt"
import "net/http"

func indextow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>American</h1>")
}

func main() {
	http.HandleFunc("/", indextow)
	fmt.Println("Start...")
	http.ListenAndServe(":2200", nil)
}
