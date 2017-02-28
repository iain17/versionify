package latest

import (
	"fmt"
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
	"net/http"
)

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the middleware :D")
		next.ServeHTTP(w, r)
	})
}

var Version *versionify.Version

func Setup(vy versionify.Versionify) {
	var err error
	Version, err = vy.NewVersion("3.0")
	if err != nil {
		panic(err)
	}

	http_versionify.NewHandlerFunc(Version, "/foo", foo, "GET").Use(exampleMiddleware)
	http_versionify.NewHandlerFunc(Version, "/bar", bar, "GET").Constrain(">= 3.0")
}
