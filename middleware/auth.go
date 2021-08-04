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

type AuthMiddleware struct {
	excluded []*controller.Route
	secret   []byte
}

func NewAuthMiddleware(secret []byte, excluded []*controller.Route) *AuthMiddleware {
	return &AuthMiddleware{
		excluded,
		secret,
	}
}

func (a *AuthMiddleware) Run(next http.Handler) http.Handler {
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

		match, _ := regexp.MatchString(`^Bearer ([\w\-]+\.){2}([\w\-]+){1}$`, token)
		if !match {
			response.Unauthorized(w, "invalid token format")
			return
		}

		t := strings.Split(token, " ")[1]
		payload := &security.Claims{}

		tk, err := jwt.ParseWithClaims(t, payload, func(jwtToken *jwt.Token) (interface{}, error) {
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
