package validator

import (
	"fmt"

	vld "github.com/go-playground/validator/v10"
)

type (
	ValidatorI interface {
		Validate(data interface{}) error
		Engine() any
	}

	ValidatorEngine struct {
		ValidatorI
		*vld.Validate
	}
)

func (v *ValidatorEngine) ValidateStruct(data interface{}) error {
	err := v.Struct(data)
	if err != nil {
		return fmt.Errorf("%s", err.(vld.ValidationErrors)[0].Error())
	}
	return nil
}

func (v *ValidatorEngine) Engine() any {
	return v.Validate
}
