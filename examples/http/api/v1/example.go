package v1

import (
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo version 1!"))
}

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 1!"))
}
