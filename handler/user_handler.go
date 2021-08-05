package handler

import (
	"encoding/json"
	"frm/model"
	"frm/repository"
	"frm/request"
	"net/http"
)

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(r *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: r}
}

func (u *UserHandler) List() (*model.Users, error) {
	return u.repo.FindAll()
}

func (u *UserHandler) CreateUser(r *http.Request) (*model.User, error) {
	req := &request.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	if ok, err := req.Validate(); !ok {
		return nil, err
	}

	user := model.NewUser(req)
	if err := u.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
