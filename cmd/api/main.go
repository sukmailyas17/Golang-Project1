package main

import (
	"fmt"
	"log"
	"net/http"
)

//this is the main project
const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

	// read from command line

	// connect to database

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}

//apa
