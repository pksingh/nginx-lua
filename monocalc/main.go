package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	AppVersion  = "v1"
	AppBasepath = "/api/" + AppVersion
	AppService  = "name: monocalc, version: " + AppVersion
)

func main() {
	handle("/status", http.MethodGet, func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"data":"ok"}`))
	})

	handle("/calculate", http.MethodPost, func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"data":"called /calculate"}`))
	})

	log.Println("Server Starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))

}

// Will use for api path and version
func prefix(route string) string {
	return fmt.Sprintf("%s%s", AppBasepath, route)
}

// handle() will add required prefix
func handle(route string, method string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(prefix(route), func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		handler(rw, req)
	})
}
