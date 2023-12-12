package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running API at port :5000")
	router := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", router))
}