package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var numberOfAccesses int64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&numberOfAccesses, 1)
		_, err := w.Write([]byte(fmt.Sprintf("Número de acessos: %d", numberOfAccesses)))
		if err != nil {
			log.Println("[main] Erro ao escrever na resposta:", err)
			return
		}
	})
	log.Println("[main] Subindo servidor...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("[main] Erro ao subir o servidor:", err)
		return
	}

}
