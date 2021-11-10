package go_learn_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(rw, r, "./resources/ok.html")
	} else {
		http.ServeFile(rw, r, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourcesOk string

//go:embed resources/notfound.html
var resourcesNotFound string

func ServeFileEmbed(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") != "" {
		fmt.Fprint(rw, resourcesOk)
	} else {
		fmt.Fprint(rw, resourcesNotFound)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
