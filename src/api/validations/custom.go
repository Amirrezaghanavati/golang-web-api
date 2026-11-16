package validations

import (
	"errors"
	"reflect"

	"github.com/Amirrezaghanavati/golang-web-api/common"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil
}

func IranianMobileNumberValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	return ok && common.IsIranianMobileNumber(value)
}

func PasswordValidator(policy *common.PasswordPolicy) validator.Func {
	return func(field validator.FieldLevel) bool {
		value := field.Field()
		return value.Kind() == reflect.String && policy.CheckPassword(value.String())
	}
}
