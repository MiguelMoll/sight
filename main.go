package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("in handle func 1")
		io.WriteString(w, "Hello from a NEW HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("in handle func 2")
		io.WriteString(w, "Hello from a NEW HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	fmt.Printf("listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
