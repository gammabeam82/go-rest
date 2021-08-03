package router

import (
	"frm/middleware"
	"frm/response"
	"github.com/gorilla/mux"
	"log"
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

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found:", r.RequestURI, r.Method)
		response.NotFound(w, "path does not exist")
	})

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
