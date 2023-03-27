package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
)

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

	// Get a client to the Reverser component.
	mathserve, err := weaver.Get[Mathserve](root)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		num1, _ := strconv.Atoi(r.URL.Query().Get("num1"))
		num2, _ := strconv.Atoi(r.URL.Query().Get("num2"))
		num, err := mathserve.Add(r.Context(), num1, num2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "The answer is, %d!\n", num)
	})

	http.HandleFunc("/sub", func(w http.ResponseWriter, r *http.Request) {
		num1, _ := strconv.Atoi(r.URL.Query().Get("num1"))
		num2, _ := strconv.Atoi(r.URL.Query().Get("num2"))
		num, err := mathserve.Sub(r.Context(), num1, num2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "The answer is, %d!\n", num)
	})

	http.Serve(lis, nil)
}
