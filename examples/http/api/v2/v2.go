package v2

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
)

func Setup(vy versionify.Versionify) {
	v, err := vy.NewVersion("2.0")
	if err != nil {
		panic(err)
	}
	//Routes
	http_versionify.NewHandlerFunc(v, "/bar", bar, "GET")
	http_versionify.NewHandlerFunc(v, "/iain", iain, "GET").Constrain("<= 2.0")
}
