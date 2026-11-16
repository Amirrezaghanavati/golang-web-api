package validations

import (
	"reflect"

	"github.com/Amirrezaghanavati/golang-web-api/common"
	"github.com/go-playground/validator/v10"
)

func PasswordValidator(policy *common.PasswordPolicy) validator.Func {
	return func(field validator.FieldLevel) bool {
		value := field.Field()
		return value.Kind() == reflect.String && policy.CheckPassword(value.String())
	}
}
