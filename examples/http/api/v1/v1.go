package v1

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

func Setup(vy versionify.Versionify) {
	v, err := vy.NewVersion("1.0")
	if err != nil {
		panic(err)
	}

	http_versionify.NewHandlerFunc(v, "/foo", foo, "GET").Use(exampleMiddleware)
	http_versionify.NewHandlerFunc(v, "/bar", bar, "GET")
}
