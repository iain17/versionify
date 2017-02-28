package api

import (
	"fmt"
	"github.com/Pocketbrain/versionify"
	"github.com/Pocketbrain/versionify/methods/http_versionify"
	logger "github.com/Sirupsen/logrus"
	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"net/http"
	"github.com/Pocketbrain/versionify/examples/http_reverse/api/v1"
	"github.com/Pocketbrain/versionify/examples/http_reverse/api/v2"
	"github.com/Pocketbrain/versionify/examples/http_reverse/api/latest"
)


func Setup() {
	//Initialize versionify
	vy := versionify.New()
	vy.SetReverse(true)

	//Call the setup methods per version
	v1.Setup(vy)
	v2.Setup(vy)
	latest.Setup(vy)

	//Connect versions to router
	router := mux.NewRouter()

	//Add any global middleware to the router
	vy.Register(http_versionify.NewMuxRegistrator(router))

	//Print the routes
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	//Default simple negroni code.
	n := negroni.Classic()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(router)
	logger.Info("Api server running on :8080")
	logger.Error(http.ListenAndServe(":8080", n))
}
