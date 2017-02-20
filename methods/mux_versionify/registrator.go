package mux_versionify

import (
	"github.com/Pocketbrain/versionify"
	"github.com/gorilla/mux"
	"fmt"
)

func NewRegistrator(router *mux.Router) versionify.VersionRegistrator {
	//Called for each version.
	return func (version *versionify.Version, methods versionify.Methods) {
		path := "/v"+version.String()
		fmt.Printf("Registering sub route '%s'\n", path)
		subRouter := router.PathPrefix(path).Subrouter()
		for name, _method := range methods {
			fmt.Printf("Registering handler '%s'\n", name)
			_method.(*MuxMethod).registerFunc(subRouter)
		}
	}
}