package controller

import (
	"frm/response"
	"net/http"
)

type IndexController struct {
	routes []*Route
}

func (i *IndexController) Routes() []*Route {
	return i.routes
}

func NewIndexController() *IndexController {
	index := &Route{
		Path:     "/api/index",
		Method:   http.MethodGet,
		Security: false,
		Action: func(w http.ResponseWriter, r *http.Request) {
			response.NewResponse("hello", nil).Send(w, http.StatusOK)
		},
	}

	secured := &Route{
		Path:     "/api/secret",
		Method:   http.MethodGet,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			response.NewResponse("how did you get here?", nil).Send(w, http.StatusOK)
		},
	}

	return &IndexController{
		[]*Route{
			index,
			secured,
		},
	}
}
