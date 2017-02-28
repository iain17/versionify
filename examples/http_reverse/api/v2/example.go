package v2

import (
	"net/http"
)

func bjorn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bjorn on version 2! Not available on version 1!"))
}
