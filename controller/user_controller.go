package controller

import (
	"frm/handler"
	"frm/response"
	"net/http"
)

type UserController struct {
	routes  []*Route
	handler *handler.UserHandler
}

func (i *UserController) Routes() []*Route {
	return i.routes
}

func NewUserController(handler *handler.UserHandler) *UserController {
	list := &Route{
		Path:     "/api/users",
		Method:   http.MethodGet,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			users, err := handler.List()

			if err != nil {

			}

			response.NewResponse(users, nil).Send(w, http.StatusOK)
		},
	}

	return &UserController{
		[]*Route{
			list,
		},
		handler,
	}
}
