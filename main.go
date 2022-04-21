package main

import (
	"log"
	"net/http"
	"unattended-test/server"
)

func main() {
	s := server.Server{}
	router := s.Setup()

	log.Fatal(http.ListenAndServe(":8080", router))
}
