package security

import (
	"encoding/json"
	"errors"
	"frm/config"
	"frm/repository"
	"golang.org/x/crypto/bcrypt"
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
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cr.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return NewToken(user.ID, user.Role, a.config)
}
