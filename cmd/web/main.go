package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
