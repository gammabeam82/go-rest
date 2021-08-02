package request

import (
	"errors"
	"regexp"
)

type CreateUserRequest struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	RepeatedPassword string `json:"repeated_password"`
}

func (c *CreateUserRequest) Validate() (bool, error) {
	if match, _ := regexp.MatchString(`(?i)^[a-z]{2,24}$`, c.Username); !match {
		return false, errors.New("invalid username")
	}

	if match, _ := regexp.MatchString(`(?i)^[\w.\-]+@[\w.\-]{2,}.[a-z]{2,8}$`, c.Email); !match {
		return false, errors.New("invalid email")
	}

	if match, _ := regexp.MatchString(`^[\w_.\-#$]{6,24}$`, c.Password); !match {
		return false, errors.New("invalid password")
	}

	if c.Password != c.RepeatedPassword {
		return false, errors.New("passwords are not equal")
	}

	return true, nil
}

type UpdateUserRequest struct {
	Username string `json:"username"`
}

func (u *UpdateUserRequest) Validate() (bool, error) {
	if match, _ := regexp.MatchString(`(?i)^[a-z]{0,24}$`, u.Username); !match {
		return false, errors.New("invalid username")
	}

	return true, nil
}
