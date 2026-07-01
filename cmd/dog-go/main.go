package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/xiaotukema/dog-go/internal/dog"
)

func main() {
	serve := flag.Bool("serve", false, "start the HTTP API")
	addr := flag.String("addr", ":8080", "HTTP listen address")
	name := flag.String("name", "", "find a dog by name")
	flag.Parse()

	repo := dog.NewRepository()

	if *serve {
		if err := runServer(*addr, repo); err != nil {
			log.Fatal(err)
		}
		return
	}

	if *name != "" {
		found, ok := repo.FindByName(*name)
		if !ok {
			fmt.Fprintf(os.Stderr, "dog %q not found\n", *name)
			os.Exit(1)
		}
		printDog(found)
		return
	}

	printDog(repo.Random())
}

func runServer(addr string, repo *dog.Repository) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.HandleFunc("GET /dogs", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, repo.All())
	})
	mux.HandleFunc("GET /dogs/random", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, repo.Random())
	})
	mux.HandleFunc("GET /dogs/{name}", func(w http.ResponseWriter, r *http.Request) {
		found, ok := repo.FindByName(r.PathValue("name"))
		if !ok {
			http.NotFound(w, r)
			return
		}
		writeJSON(w, found)
	})

	log.Printf("dog-go listening on %s", addr)
	return http.ListenAndServe(addr, mux)
}

func printDog(d dog.Dog) {
	fmt.Printf("%s is a %d-year-old %s who is %s.\n", d.Name, d.Age, d.Breed, strings.Join(d.Personality, ", "))
}

func writeJSON(w http.ResponseWriter, value any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
