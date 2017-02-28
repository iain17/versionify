package handlers

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
	"net/http"
)

func V3(v versionify.Versionify) {
	v3, err := v.NewVersion("3.0")
	if err != nil {
		panic(err)
	}
	//Routes
	http_versionify.NewHandlerFunc(v3, "/bjorn", bjornV3, "", "GET")
}

func bjornV3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bjorn on version 3! Not available on version 2!"))
}
