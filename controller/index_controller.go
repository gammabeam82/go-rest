package controller

import (
	"frm/response"
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
			response.NewResponse("hello", nil).Send(w, http.StatusOK)
		},
	}

	return &IndexController{
		[]*router.Route{
			index,
		},
	}
}
