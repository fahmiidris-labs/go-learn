package go_learn_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(rw, r)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(rw, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(rw, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(rw, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		fmt.Fprint(rw, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Foo Executed")
		panic("Ups")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
