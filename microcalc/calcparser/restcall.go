package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func DoPost(operator string, opr, opl int) (*http.Response, error) {
	requestURL := NodeEndpoint[operator]
	jsonBody := fmt.Sprintf(`{ "operands": [ %d, %d ] }`, opr, opl)
	bodyReader := bytes.NewReader([]byte(jsonBody))
	return res, nil
}
