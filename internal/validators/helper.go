package validators

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateRequest(req interface{}) error {
	validate := getDecortValidator()
	return validate.Struct(req)
}

func ValidationError(fe validator.FieldError) error {
	return errors.New(errorMessage(fe))
}

//nolint:errorlint
func GetErrors(err error) validator.ValidationErrors {
	return err.(validator.ValidationErrors)
}

func StringInSlice(str string, target []string) bool {
	for _, v := range target {
		if v == str {
			return true
		}
	}
	return false
}
