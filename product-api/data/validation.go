package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"key: '%s' error: field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors converts the slice into a string slice
func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains
type Validation struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validation{validate}
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
//	if ve, ok := err.(validator.ValidationErrors); ok {
//				fmt.Println(ve.Namespace())
//				fmt.Println(ve.Field())
//				fmt.Println(ve.StructNamespace())
//				fmt.Println(ve.StructField())
//				fmt.Println(ve.Tag())
//				fmt.Println(ve.ActualTag())
//				fmt.Println(ve.Kind())
//				fmt.Println(ve.Type())
//				fmt.Println(ve.Value())
//				fmt.Println(ve.Param())
//				fmt.Println()
//		}
func (v *Validation) Validate(i interface{}) ValidationErrors {
	// errs := v.validate.Struct(i).(validator.ValidationErrors)
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}

	errors := errs.(validator.ValidationErrors)

	if len(errors) == 0 {
		return nil
	}

	var returnErrs []ValidationError
	for _, err := range errors {
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}

	return returnErrs
}

// validateSKU
func validateSKU(fl validator.FieldLevel) bool {
	// SKU must be in the format abc-abc-abc
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	sku := re.FindAllString(fl.Field().String(), -1)
	fmt.Println("LOGGG VALIDATION 1, len of sku: ", len(sku))

	return len(sku) == 1
}
