package handlers

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/mux_versionify"
	"net/http"
)

func V2(v versionify.Versionify) {
	v2, err := v.NewVersion("2.0")
	if err != nil {
		panic(err)
	}
	//Routes
	mux_versionify.NewHandlerFunc(v2, "/bar", barV2, "", "GET")
	mux_versionify.NewHandlerFunc(v2, "/iain", iainV2, "<= 2.0", "GET")
}

func barV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 2!"))
}

func iainV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Iain on version 2! Not available on version 1!"))
}
