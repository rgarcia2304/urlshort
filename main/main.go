package main

import (
	"fmt"
	"net/http"
	"github.com/rgarcia2304/urlshort"
	"log"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlData := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	jsonData :=
	`[
	{"path": "/short", "url": "https://github.com/rgarcia2304"},
	{"path": "/big",    "url": "https://youtube.com"}
]`
	jsonHandler, err := urlshort.JSONHandler([]byte(jsonData),mapHandler)
	yamlHandler, err := urlshort.YAMLHandler([]byte(yamlData), jsonHandler)

	log.Printf("yamlHandler is nil? %v", yamlHandler == nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080",yamlHandler)
	
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
