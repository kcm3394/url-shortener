package main

import (
	"flag"
	"fmt"
	"github.com/kcm3394/url-shortener/urlshort"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var (
	fileName = flag.String("file", "default.yml", "the file containing shortened paths to URLs in either yaml or json format")
)

func main() {
	mux := defaultMux()

	flag.Parse()

	ext := filepath.Ext(*fileName)

	var handler http.Handler
	var err error
	if ext == ".yml" {
		handler, err = urlshort.YAMLHandler(getBytesFromFile(*fileName), mux)
		if err != nil {
			panic(err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(getBytesFromFile(*fileName), mux)
		if err != nil {
			panic(err)
		}
	} else {
		log.Fatalf("Unsupported file type: %s", ext)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getBytesFromFile(fileName string) []byte {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Cannot read file %s: %v", fileName, err)
	}
	return f
}
