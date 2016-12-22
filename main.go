package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=====================")
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Path: %s\n", r.URL.Path)
	fmt.Printf("Query: %+v\n", r.URL.Query())
	body, err := getRequestBody(r)
	if err != nil {
		fmt.Println("Error getting request body:")
		fmt.Println(err)
	}
	fmt.Printf("Body: %s\n", body)
	fmt.Printf("Headers: %+v\n", r.Header)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("App running")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getRequestBody(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	// Restore r.Body after ioutil.ReadAll
	// https://medium.com/@xoen/golang-read-from-an-io-readwriter-without-loosing-its-content-2c6911805361
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return string(body), nil
}
