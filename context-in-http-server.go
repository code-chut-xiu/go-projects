package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Server: hello handler started")
	defer fmt.Println("Server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello from Go!\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Server: ", err)
		internalErr := http.StatusInternalServerError
		http.Error(w, err.Error(), internalErr)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}