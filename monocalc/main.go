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
	http.HandleFunc("/status", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"data":"ok"}`))
	})

	log.Println("Server Starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))

}
