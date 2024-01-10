package main

import (
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./24-FileServer/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<form method="POST" action="/process">
				<input type="text" name="name">
				<input type="submit" value="Submit">
			</form>
		`))
	})

	http.ListenAndServe(":8081", mux)
}
