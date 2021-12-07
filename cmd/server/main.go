package main

import (
	"log"

	"github.com/hokdre/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8081")
	log.Fatal(srv.ListenAndServe())
}
