package data

import (
	"fmt"

	"github.com/go-playground/validator"
)

// ValidationError is wrapper for validators FieldError
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// Collection of ValidationError
type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()

	return &Validation{validate}
}

// Validate the item. Returned errors are
// cast into validator.ValidationErrors
func (v *Validation) Validate(i interface{}) ValidationErrors {
	// Validate the item
	errs := v.validate.Struct(i)

	if errs == nil {
		return nil
	}

	castedErrs := errs.(validator.ValidationErrors)
	var returnErrs ValidationErrors
	for _, err := range castedErrs {
		// Cast FieldErrors into ValidationError
		ve := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}
