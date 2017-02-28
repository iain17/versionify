package handlers

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
	"net/http"
)

func V2(v versionify.Versionify) {
	v2, err := v.NewVersion("2.0")
	if err != nil {
		panic(err)
	}
	//Routes
	http_versionify.NewHandlerFunc(v2, "/bar", barV2, "GET")
	http_versionify.NewHandlerFunc(v2, "/iain", iainV2, "GET").Constrain("<= 2.0")
}

func barV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 2!"))
}

func iainV2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Iain on version 2! Not available on version 1!"))
}
