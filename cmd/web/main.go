package main

import (
	"flag"
	"log"
	"net/http"
)

//TODO 1
//1) add INFO and ErrorLog
//2) Create struct application
//3) DI INFO and ErrorLog
//4) Centralize Error Handling (touch helpers.go and add methods to application struct
//	 for common err (server, client, notfound))
//
func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting  server on %v", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		panic(err)
	}
}
