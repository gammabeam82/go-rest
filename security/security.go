package security

import (
	"frm/config"
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

func NewToken(userId uint, role string, config *config.Config) (*Token, error) {
	claims := Claims{
		UserId: userId,
		Role:   role,
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
		UserId: userId,
		Token:  "Bearer " + ts,
	}, nil
}
