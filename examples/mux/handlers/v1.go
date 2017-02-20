package handlers

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/mux_versionify"
	"net/http"
)

func V1(v versionify.Versionify) {
	v1, err := v.NewVersion("1.0")
	if err != nil {
		panic(err)
	}
	//Routes
	mux_versionify.NewRoute(v1, "/foo", fooV1, "", "GET")
	mux_versionify.NewRoute(v1, "/bar", barV1, "", "GET")
}

func fooV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo version 1!"))
}

func barV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 1!"))
}