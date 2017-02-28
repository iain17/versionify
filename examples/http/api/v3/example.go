package v3

import (
	"net/http"
)

func bjorn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bjorn on version 3! Not available on version 2!"))
}
