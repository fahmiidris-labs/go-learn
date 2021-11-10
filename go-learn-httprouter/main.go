package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(rw, "Hello HttpRouter")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	server.ListenAndServe()
}
