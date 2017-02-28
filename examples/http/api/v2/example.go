package v2

import (
	"net/http"
)

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 2!"))
}

func iain(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Iain on version 2! Not available on version 1!"))
}
