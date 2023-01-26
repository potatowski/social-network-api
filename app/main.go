package main

import (
	"fmt"
	"log"
	"net/http"
	"social-api/src/router"
)

func main() {
	router := router.Generate()
	fmt.Println("Routing API")
	log.Fatal(http.ListenAndServe((":5000"), router))
}
