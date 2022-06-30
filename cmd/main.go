package main

import (
	"log"
	"msqrd/pkg/proxy"
	"net/http"
)

func main() {
	log.Println("start proxy server")
	if err := http.ListenAndServe(":8080", new(proxy.Proxy)); err != nil {
		log.Fatal(err)
	}
}
