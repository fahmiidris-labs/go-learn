package go_learn_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "name is empty")
	} else {
		fmt.Fprintf(rw, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}
