package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func NewValidator() *validator.Validate {
	v := validator.New()

	_ = v.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`(?i)^[a-z]{2,24}$`, fl.Field().String())

		return match
	})

	_ = v.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		match, _ := regexp.MatchString(`^[\w_.\-#$]{6,24}$`, fl.Field().String())

		return match
	})

	return v
}
