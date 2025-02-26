package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := 8888
	if os.Getenv("PORT") != "" {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err == nil {
			port = p
		}
	}

	fmt.Println(fmt.Sprintf("Starting server on port %d...", port))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
