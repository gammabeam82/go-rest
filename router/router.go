package router

import (
	"frm/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type HasRoutes interface {
	Routes() []*Route
}

type Route struct {
	Path     string
	Method   string
	Security bool
	Action   http.HandlerFunc
}

func NewRouter(controllers ...HasRoutes) *mux.Router {
	var anonymous []*Route

	router := mux.NewRouter()

	for _, controller := range controllers {
		for _, route := range controller.Routes() {

			if !route.Security {
				anonymous = append(anonymous, route)
			}

			router.HandleFunc(route.Path, route.Action).Methods(route.Method)
		}
	}

	router.Use(middleware.Logger)

	return router
}
