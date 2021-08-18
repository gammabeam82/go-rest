package security

import (
	"fmt"
	"frm/config"
	"frm/model"
	"github.com/golang-jwt/jwt"
	"time"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type Claims struct {
	UserId uint
	Role   string
	jwt.StandardClaims
}

func NewToken(user *model.User, config *config.Config) (*Token, error) {
	claims := Claims{
		UserId: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.JwtTokenTTL(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ts, err := token.SignedString(config.JwtSecret())

	if err != nil {
		return nil, err
	}

	return &Token{
		UserId: user.ID,
		Token:  fmt.Sprintf("Bearer %s", ts),
	}, nil
}
