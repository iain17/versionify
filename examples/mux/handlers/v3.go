package handlers

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/mux_versionify"
	"net/http"
)

func V3(v versionify.Versionify) {
	v3, err := v.NewVersion("3.0")
	if err != nil {
		panic(err)
	}
	//Routes
	mux_versionify.NewRoute(v3, "/bjorn", bjornV3, "", "GET")
}

func bjornV3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Iain on version 2! Not available on version 1!"))
}