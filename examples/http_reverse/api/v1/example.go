package v1

import (
	"net/http"
)

func bar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 1!"))
}

func deprecated_idea(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deprecated idea on version 1! Not available on version 2!"))
}
