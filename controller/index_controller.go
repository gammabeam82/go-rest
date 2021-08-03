package controller

import (
	"frm/router"
	"net/http"
)

type IndexController struct {
	routes []*router.Route
}

func (u *IndexController) Routes() []*router.Route {
	return u.routes
}

func NewIndexController() *IndexController {
	index := &router.Route{
		Path:     "/api/index",
		Method:   http.MethodGet,
		Security: false,
		Action: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		},
	}

	return &IndexController{
		[]*router.Route{
			index,
		},
	}
}
