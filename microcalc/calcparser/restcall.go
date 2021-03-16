package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func DoPost(operator string, opr, opl int) (*http.Response, error) {
	requestURL := NodeEndpoint[operator]
	jsonBody := fmt.Sprintf(`{ "operands": [ %d, %d ] }`, opr, opl)
	bodyReader := bytes.NewReader([]byte(jsonBody))
	req, _ := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, _ := client.Do(req)
	return res, nil
}
