package handler

import (
	"encoding/json"
	"frm/model"
	"frm/repository"
	"frm/request"
	"net/http"
)

type AccessDenied struct{}

func (a AccessDenied) Error() string {
	return "access denied"
}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(r *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: r}
}

func (u *UserHandler) List() (*model.Users, error) {
	return u.repo.FindAll()
}

func (u *UserHandler) GetUser(id int) (*model.User, error) {
	return u.repo.FindById(id)
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

func (u *UserHandler) DeleteUser(currentUser *model.User, id int) error {
	user, err := u.repo.FindById(id)

	if err != nil {
		return err
	}

	if !currentUser.CanDelete(user) {
		return &AccessDenied{}
	}

	if err = u.repo.Delete(user); err != nil {
		return err
	}

	return nil
}

func (u *UserHandler) UpdateUser(r *http.Request, currentUser *model.User, id int) error {
	user, err := u.repo.FindById(id)

	if err != nil {
		return err
	}

	if !currentUser.CanUpdate(user) {
		return &AccessDenied{}
	}

	req := &request.UpdateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	if ok, err := req.Validate(); !ok {
		return err
	}

	user.Rename(req)

	return u.repo.Update(user)
}
