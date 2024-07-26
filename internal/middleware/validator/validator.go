package validator

import (
	"net/mail"
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidateRegex(regex, value string) bool {
	reg := regexp.MustCompile(regex)
	return reg.Match([]byte(value))
}

func Phone() validator.Func {
	const phoneRegex = `^01([0|1|6|7|8|9])([0-9]{3,4})([0-9]{4})$`
	return func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			return ValidateRegex(phoneRegex, value)
		}
		return false
	}
}

func Email() validator.Func {
	return func(fl validator.FieldLevel) bool {
		if value, ok := fl.Field().Interface().(string); ok {
			_, err := mail.ParseAddress(value)
			return err == nil
		}
		return false
	}
}

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", Phone())
		v.RegisterValidation("email", Email())
	}
}
