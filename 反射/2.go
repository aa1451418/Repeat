package main

import (
	"fmt"
	"net/http"
)

func indexone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello is fangshai</h1>")
}

func main() {
	http.HandleFunc("/", indexone)
	fmt.Println("Start_Service........")
	http.ListenAndServe(":24354", nil)

}
