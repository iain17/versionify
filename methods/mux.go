package methods

import (
	hversion "github.com/hashicorp/go-version"
	"github.com/gorilla/mux"
	"github.com/Pocketbrain/versionify"
	"net/http"
	"fmt"
)

//A versionify.Method implementation for the mux router. Commonly used in web packages to register the routing of a web application.
type RegisterRoutesFunc func(r *mux.Router)

type MuxMethod struct {
	Constraints hversion.Constraints
	registerFunc RegisterRoutesFunc
}

func NewMuxMethod(registerFunc RegisterRoutesFunc, constraints hversion.Constraints) versionify.Method {
	return &MuxMethod{
		registerFunc: registerFunc,
		Constraints: constraints,
	}
}

//A simplification of NewMuxMethod
//constraint ">= 1.0, < 1.4"
func NewMuxRoute(version *versionify.Version, path string, f func(http.ResponseWriter, *http.Request), constraints string, methods ...string) versionify.Method {
	name := path
	register := func(r *mux.Router) {
		r.HandleFunc(path, f).Methods(methods...)
	}
	var Constraints hversion.Constraints
	var err error
	if constraints != "" {
		Constraints, err = hversion.NewConstraint(constraints)
		if err != nil {
			fmt.Println(fmt.Errorf("Could not create constraint: %v", err))
		}
	}
	method, err := version.Method(name, NewMuxMethod(register, Constraints))
	if err != nil {
		panic(fmt.Errorf("Could not create mux route: %v", err))
	}
	return method
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