package validation

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(data interface{}) error {
	errs := validate.Struct(data)
	if errs != nil {
		return errs
	}

	return nil
}
