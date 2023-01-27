package main

import (
	"fmt"
	"log"
	"net/http"
	"social-api/src/config"
	"social-api/src/router"
)

func main() {
	config.Load()
	router := router.Generate()
	fmt.Printf("Routing API in port %d!", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
