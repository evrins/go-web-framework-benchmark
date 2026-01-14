package server

import (
	"fmt"
	"log"
	"net/http"
)

func listenAndServe(port int, handler http.Handler) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
