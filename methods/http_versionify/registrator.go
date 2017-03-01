package http_versionify

import (
	"github.com/Pocketbrain/versionify"
	"github.com/gorilla/mux"
)

//Easy out of the box registrator for the mux router.
//Clone this for any other router you might have.
func NewMuxRegistrator(router *mux.Router) versionify.VersionRegistrator {
	//Called for each version.
	return func(version *versionify.Version, methods versionify.Methods) {
		path := "/v" + version.String()
		//fmt.Printf("Registering sub route '%s'\n", path)
		versionRouter := router.PathPrefix(path).Subrouter()
		for _, method := range methods {
			if  httpMethod, ok := method.(*HttpMethod); ok {
				handler := httpMethod.getHandler()
				versionRouter.Handle(httpMethod.path, handler).Methods(httpMethod.methods...)
				//fmt.Printf("Registering handler with name '%s' and path %s\n", name, path + httpMethod.path)
			}
		}
	}
}
