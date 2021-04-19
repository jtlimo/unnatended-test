package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server running at http://localhost:9091")
	fmt.Println("Ctrl+C or Ctrl+D to quit")
	log.Fatal((http.ListenAndServe(":9091", nil)))
}
