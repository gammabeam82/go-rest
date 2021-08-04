package controller

import (
	"frm/response"
	"frm/security"
	"net/http"
)

type SecurityController struct {
	authenticator *security.Authenticator
	routes        []*Route
}

func (s *SecurityController) Routes() []*Route {
	return s.routes
}

func NewSecurityController(authenticator *security.Authenticator) *SecurityController {
	login := &Route{
		Path:     "/api/login",
		Method:   http.MethodPost,
		Security: false,
		Action: func(w http.ResponseWriter, r *http.Request) {
			token, err := authenticator.Login(r)

			if err != nil {
				response.NewResponse(nil, err).Send(w, http.StatusBadRequest)
				return
			}

			response.NewResponse(token, nil).Send(w, http.StatusOK)
		},
	}

	return &SecurityController{
		authenticator,
		[]*Route{
			login,
		},
	}
}
