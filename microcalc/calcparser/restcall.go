package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

func DoPost(operator string, opr, opl int) (*http.Response, error) {
	requestURL := NodeEndpoint[operator]
	jsonBody := fmt.Sprintf(`{ "operands": [ %d, %d ] }`, opr, opl)
	bodyReader := bytes.NewReader([]byte(jsonBody))
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	return res, nil
}
