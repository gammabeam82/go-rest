package router

import (
	"frm/config"
	"frm/controller"
	"frm/middleware"
	"frm/response"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter(config *config.Config, controllers ...controller.HasRoutes) *mux.Router {
	var excluded []*controller.Route

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Not found:", r.RequestURI, r.Method)
		response.NotFound(w, "path does not exist")
	})

	for _, c := range controllers {
		for _, route := range c.Routes() {

			if !route.Security {
				excluded = append(excluded, route)
			}

			router.HandleFunc(route.Path, route.Action).Methods(route.Method)
		}
	}

	accessManager := middleware.NewAccessManager(config.JwtSecret(), excluded)

	router.Use(middleware.Logger, accessManager.Run)

	return router
}
