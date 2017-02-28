package v1

import (
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
)

func Setup(vy versionify.Versionify) {
	v, err := vy.NewVersion("1.0")
	if err != nil {
		panic(err)
	}
	//Routes
	http_versionify.NewHandlerFunc(v, "/bar", bar, "GET")
	http_versionify.NewHandlerFunc(v, "/super_deprecated_idea", deprecated_idea, "GET").Constrain("< 2.0")
}
