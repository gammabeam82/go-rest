package security

import (
	"encoding/json"
	"frm/config"
	"frm/repository"
	"frm/service"
	"net/http"
)

type Authenticator struct {
	repo   *repository.UserRepository
	config *config.Config
}

func NewAuthenticator(c *config.Config, r *repository.UserRepository) *Authenticator {
	return &Authenticator{
		repo:   r,
		config: c,
	}
}

func (a Authenticator) Login(r *http.Request) (*Token, error) {
	cr := &Credentials{}

	if err := json.NewDecoder(r.Body).Decode(cr); err != nil {
		return nil, err
	}

	user, err := a.repo.FindByEmail(cr.Email)

	if err != nil {
		return nil, err
	}

	if err := service.IsPasswordValid(user.Password, cr.Password); err != nil {
		return nil, err
	}

	return NewToken(user.ID, user.Role, a.config)
}
