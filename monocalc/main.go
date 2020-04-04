package main

import (
	"log"
	"net/http"
)

const (
	AppVersion  = "v1"
	AppBasepath = "/api/" + AppVersion
	AppService  = "name: monocalc, version: " + AppVersion
)

func main() {
	log.Println("Server Starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))

}
