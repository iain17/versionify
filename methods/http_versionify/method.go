package http_versionify

import (
	"fmt"
	"github.com/Pocketbrain/versionify"
	hversion "github.com/hashicorp/go-version"
	"net/http"
	"strings"
)

//A versionify.Method implementation of http.
type HttpMethod struct {
	constraints hversion.Constraints
	path        string
	handler     http.Handler
	methods     []string
	middleware  []func(http.Handler) http.Handler
}

// Handler sets a handler for the route.
func NewHandler(version *versionify.Version, path string, handler http.Handler, methods ...string) *HttpMethod {
	name := strings.Join(methods, "_") + path
	muxMethod := &HttpMethod{
		path:       path,
		handler:    handler,
		methods:    methods,
		middleware: []func(http.Handler) http.Handler{},
	}
	_, err := version.Method(name, muxMethod)
	if err != nil {
		panic(fmt.Errorf("Could not create mux route: %v", err))
	}
	return muxMethod
}

// HandlerFunc sets a handler function for the route.
func NewHandlerFunc(version *versionify.Version, path string, f func(http.ResponseWriter, *http.Request), methods ...string) *HttpMethod {
	return NewHandler(version, path, http.HandlerFunc(f), methods...)
}

//Add middleware to the route.
func (m *HttpMethod) Use(middleware ...func(http.Handler) http.Handler) *HttpMethod {
	m.middleware = middleware
	return m
}

// Middleware chainer. Returns the handler.
func (m *HttpMethod) getHandler() http.Handler {
	h := m.handler
	for _, m := range m.middleware {
		h = m(h)
	}
	return h
}

/**
Add constraints. Defining for which version a method is final deprecated.
*/
func (m *HttpMethod) Constrain(constraints string) *HttpMethod {
	var err error
	m.constraints, err = hversion.NewConstraint(constraints)
	if err != nil {
		fmt.Println(fmt.Errorf("Could not create constraint: %v", err))
	}
	return m
}

//TODO add deprecated warning middleware.

/**
Called when a method is added for a version, to check a possible set constraint.
*/
func (m *HttpMethod) Check(v *versionify.Version) bool {
	if m.constraints == nil {
		return true
	}
	return m.constraints.Check(&v.Version)
}
