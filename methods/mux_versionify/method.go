package mux_versionify

import (
	"fmt"
	"github.com/Pocketbrain/versionify"
	"github.com/gorilla/mux"
	hversion "github.com/hashicorp/go-version"
	"net/http"
)

//A versionify.Method implementation for the mux router. Commonly used in web packages to register the routing of a web application.
type RegisterRoutesFunc func(r *mux.Router)

type MuxMethod struct {
	Constraints  hversion.Constraints
	registerFunc RegisterRoutesFunc
}

func NewMethod(registerFunc RegisterRoutesFunc, constraints hversion.Constraints) versionify.Method {
	return &MuxMethod{
		registerFunc: registerFunc,
		Constraints:  constraints,
	}
}

// Handler sets a handler for the route.
func NewHandler(version *versionify.Version, path string, handler http.Handler, constraints string, methods ...string) versionify.Method {
	name := path
	register := func(r *mux.Router) {
		r.PathPrefix(path).Handler(handler).Methods(methods...)
	}
	var Constraints hversion.Constraints
	var err error
	if constraints != "" {
		Constraints, err = hversion.NewConstraint(constraints)
		if err != nil {
			fmt.Println(fmt.Errorf("Could not create constraint: %v", err))
		}
	}
	method, err := version.Method(name, NewMethod(register, Constraints))
	if err != nil {
		panic(fmt.Errorf("Could not create mux route: %v", err))
	}
	return method
}

// HandlerFunc sets a handler function for the route.
func NewHandlerFunc(version *versionify.Version, path string, f func(http.ResponseWriter, *http.Request), constraints string, methods ...string) versionify.Method {
	return NewHandler(version, path, http.HandlerFunc(f), constraints, methods...)
}

/**
Called when a method is added for a version, to check a possible set constraint.
*/
func (m *MuxMethod) Check(v *versionify.Version) bool {
	if m.Constraints == nil {
		return true
	}
	return m.Constraints.Check(&v.Version)
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		},
	)
}
