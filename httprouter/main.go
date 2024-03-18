package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	// Automatic OPTIONS
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
			header.Set("Access-Control-Max-Age", "3600")
		}
		header.Set("hogehoge", "piyo")
		w.WriteHeader(http.StatusNoContent)
	})
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
