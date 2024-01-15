package main

import (
	"fmt"
	"net/http"
	"time"
)

var number uint64 = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante numero %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

//teste no apache benchmark
//ab -n 10000 -c 100 http://localhost:3000/
