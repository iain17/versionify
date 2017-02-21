package main

import (
	"fmt"
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/examples/mux/handlers"
	"github.com/Pocketbrain/versionify/methods/mux_versionify"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

func main() {
	//Initialize versionify
	v := versionify.New()

	//Registers our versions.
	handlers.V1(v)
	handlers.V2(v)
	handlers.V3(v)

	//Connect versions to router
	router := mux.NewRouter()
	v.Register(mux_versionify.NewRegistrator(router))

	//Default simple negroni code.
	n := negroni.Classic()
	n.UseHandler(router)
	fmt.Println(http.ListenAndServe(":8080", n))
}
