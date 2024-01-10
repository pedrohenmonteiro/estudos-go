package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.Handle("/blog", blog{title: "bloog"})
	http.ListenAndServe(":8080", mux)
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
