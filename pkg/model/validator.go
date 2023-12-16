package model

import (
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

const (
	StructValidatorTagName = "binding"
)

type structValidator struct {
	*validator.Validate
}

func (v *structValidator) Engine() interface{} {
	return nil
}

func (v *structValidator) ValidateStruct(i interface{}) error {
	return v.Struct(i)
}

// New creates a new struct validator
func New() binding.StructValidator {
	v := validator.New()
	v.SetTagName(StructValidatorTagName)

	// Register validators here

	return &structValidator{v}
}
