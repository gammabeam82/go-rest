package controller

import "net/http"

type HasRoutes interface {
	Routes() []*Route
}

type Route struct {
	Path     string
	Method   string
	Security bool
	Action   http.HandlerFunc
}
