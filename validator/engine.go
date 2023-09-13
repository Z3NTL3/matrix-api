package validator

import (
	"fmt"

	vld "github.com/go-playground/validator/v10"
)

var Validator *vld.Validate

type (
	ValidatorI interface {
		Validate(data interface{}) error
	}

	ValidatorEngine struct {
		ValidatorI
		*vld.Validate
	}
)

func (v *ValidatorEngine) Validify(data interface{}) error {
	err := v.Struct(data)
	if err != nil {
		return fmt.Errorf("Did not match POSTAL CODE ZIP FORMAT [%s]", err.(vld.ValidationErrors)[0].Tag())
	}
	return nil
}
