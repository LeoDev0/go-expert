package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	startTime := time.Now()
	request, err := http.Get("http://ident.me")
	fmt.Printf("Execution time: %s\n", time.Since(startTime))
	if err != nil {
		panic(err)
	}
	defer request.Body.Close() // Close the connection after the function ends. Defers are executed in LIFO order

	response, err := io.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	println("My IP address: " + string(response))
}