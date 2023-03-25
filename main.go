package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

func Add(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	result := 0

	for i := 0; i < x; i++ {
		result += y
	}

	return result
}

func main() {
	// Get a network listener on address "localhost:12345".
	root := weaver.Init(context.Background())
	opts := weaver.ListenerOptions{LocalAddress: "localhost:12345"}
	lis, err := root.Listener("hello", opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("hello listener available on %v\n", lis)

	// Get a client to the Reverser component.
	reverser, err := weaver.Get[Reverser](root)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
	})

	// Serve the /reverse endpoint
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})

	http.HandleFunc("/reverse_2", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})

	http.Serve(lis, nil)
}
