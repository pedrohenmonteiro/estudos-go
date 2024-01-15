package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante numero %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}

//teste no apache benchmark
//ab -n 10000 -c 100 http://localhost:3000/

//verificar cenário de race condition:
//go run -race main.go
