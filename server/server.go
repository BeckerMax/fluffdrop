package server

import (
	"fluffdrop/helper"
	"fmt"
	"html"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/hellofluffy", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func upload(w http.ResponseWriter, r *http.Request) {
	storage := helper.NewlocalFileHandler()
	storage.
}

func receiveFileHandler() {
	//bla
}
