package validations

import (
	"github.com/Amirrezaghanavati/golang-web-api/common"
	"github.com/go-playground/validator/v10"
)

func IranianMobileNumberValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	return ok && common.IsIranianMobileNumber(value)
}
