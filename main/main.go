package main

import (
	"flag"
	"fmt"
	urlshort "github.com/kcm3394/url-shortener"
	"io/ioutil"
	"net/http"
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
	yamlFile := readArguments()
	yaml, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func readArguments() string {
	yamlFile := flag.String("yaml", "default.yml", "a yaml file in the format of \n'- path: /some-path\n   url: https://www.some-url.com/demo'")
	flag.Parse()

	return *yamlFile
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
