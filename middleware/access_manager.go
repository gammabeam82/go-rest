package middleware

import (
	"context"
	"frm/controller"
	"frm/model"
	"frm/response"
	"frm/security"
	"github.com/golang-jwt/jwt"
	"net/http"
	"regexp"
	"strings"
)

type AccessManager struct {
	excluded []*controller.Route
	secret   []byte
}

func NewAccessManager(secret []byte, excluded []*controller.Route) *AccessManager {
	return &AccessManager{
		excluded,
		secret,
	}
}

func (a *AccessManager) Run(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, route := range a.excluded {
			if route.Path == r.URL.Path && route.Method == r.Method {
				next.ServeHTTP(w, r)
				return
			}
		}

		token := r.Header.Get("Authorization")

		if len(token) == 0 {
			response.Unauthorized(w, "Authentication needed")
			return
		}

		match, _ := regexp.MatchString(`^Bearer ([\w\-]+\.){2}([\w\-]+)$`, token)

		if !match {
			response.Unauthorized(w, "invalid token format")
			return
		}

		payload := &security.Claims{}

		tk, err := jwt.ParseWithClaims(strings.Split(token, " ")[1], payload, func(jwtToken *jwt.Token) (interface{}, error) {
			return a.secret, nil
		})

		if err != nil || !tk.Valid {
			response.Forbidden(w, "invalid token")
			return
		}

		ctx := context.WithValue(r.Context(), "user", &model.User{
			ID:   payload.UserId,
			Role: payload.Role,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
