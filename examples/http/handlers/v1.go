package handlers

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

func V1(v versionify.Versionify) {
	v1, err := v.NewVersion("1.0")
	if err != nil {
		panic(err)
	}

	//Routes
	http_versionify.NewHandlerFunc(v1, "/foo", fooV1, "GET").Use(exampleMiddleware)
	http_versionify.NewHandlerFunc(v1, "/bar", barV1, "GET")
}

func fooV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo version 1!"))
}

func barV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar version 1!"))
}
