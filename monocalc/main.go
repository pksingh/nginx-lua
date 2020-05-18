package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"engine" // will refer to github.com/pksingh/uservices/engine
)

const (
	AppVersion  = "v1"
	AppBasepath = "/api/" + AppVersion
	AppService  = "name: monocalc, version: " + AppVersion
)

type XInt int

// Struct for input
type InputRequest struct {
	Input string `json:"input" binding:"required"`
}

func main() {
	handle("/status", http.MethodGet, func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"data":"ok"}`))
	})

	handle("/calculate", http.MethodPost, func(rw http.ResponseWriter, req *http.Request) {
		var body InputRequest

		if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Sprintf(`{"error":"Invalid Input","service":"%s"}`, AppService)))
			return
		}
		fmt.Println("Received on /calculate:", body.Input)

		infix := engine.Tokenize(body.Input)
		tokens := engine.PostFixExpression(infix)
		result := engine.Calculate(tokens)
		// result := "WIP"

		ret := map[string]interface{}{
			"operands": "[" + body.Input + "]",
			"origins":  body.Input,
			"result":   result,
			"service":  AppService,
		}
		fmt.Println("Returened on /calculate:", ret)

		if err := json.NewEncoder(rw).Encode(ret); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Println("Server Starting on 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))

}

// Will use for api path and version
func prefix(route string) string {
	return fmt.Sprintf("%s%s", AppBasepath, route)
}

// handle() will add required prefix and inject header if missed.
func handle(route string, method string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(prefix(route), func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != method {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		rw.Header().Set("Content-Type", "application/json")
		handler(rw, req)
	})
}
