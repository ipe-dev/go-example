package main

import (
	"fmt"
	"net/http"
)

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle func")
}
func main() {
	http.HandleFunc("/HandleFunc", hoge)
	http.Handle("/Handle", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "handle")
	}))
	http.ListenAndServe(":8080", nil)
}
