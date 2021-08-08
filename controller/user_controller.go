package controller

import (
	"errors"
	"frm/handler"
	"frm/model"
	"frm/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getCurrentUser(r *http.Request) (*model.User, error) {
	user, ok := r.Context().Value("user").(*model.User)

	if !ok {
		return nil, errors.New("")
	}

	return user, nil
}

type UserController struct {
	routes  []*Route
	handler *handler.UserHandler
}

func (i *UserController) Routes() []*Route {
	return i.routes
}

func NewUserController(handler *handler.UserHandler) *UserController {
	create := &Route{
		Path:     "/api/users",
		Method:   http.MethodPost,
		Security: false,
		Action: func(w http.ResponseWriter, r *http.Request) {
			user, err := handler.CreateUser(r)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			response.NewResponse(user, nil).Send(w, http.StatusCreated)
		},
	}

	list := &Route{
		Path:     "/api/users",
		Method:   http.MethodGet,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			users, err := handler.List()

			if err != nil {
				response.InternalError(w, "")
				return
			}

			response.NewResponse(users, nil).Send(w, http.StatusOK)
		},
	}

	show := &Route{
		Path:     "/api/users/{id:[0-9]+}",
		Method:   http.MethodGet,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(mux.Vars(r)["id"])

			user, err := handler.GetUser(id)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			response.NewResponse(user, nil).Send(w, http.StatusOK)
		},
	}

	update := &Route{
		Path:     "/api/users/{id:[0-9]+}",
		Method:   http.MethodPatch,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(mux.Vars(r)["id"])

			currentUser, err := getCurrentUser(r)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			err = handler.UpdateUser(r, currentUser, id)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			response.NewResponse("user was successfully updated", nil).Send(w, http.StatusOK)
		},
	}

	remove := &Route{
		Path:     "/api/users/{id:[0-9]+}",
		Method:   http.MethodDelete,
		Security: true,
		Action: func(w http.ResponseWriter, r *http.Request) {
			id, _ := strconv.Atoi(mux.Vars(r)["id"])

			currentUser, err := getCurrentUser(r)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			err = handler.DeleteUser(currentUser, id)

			if err != nil {
				response.InternalError(w, "")
				return
			}

			response.NewResponse("user was successfully removed", nil).Send(w, http.StatusOK)
		},
	}

	return &UserController{
		[]*Route{
			list,
			create,
			show,
			update,
			remove,
		},
		handler,
	}
}
