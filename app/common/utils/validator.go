package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	regular := "^1\\d{10}$"
	reg := regexp.MustCompile(regular)

	return reg.MatchString(mobile)
}
