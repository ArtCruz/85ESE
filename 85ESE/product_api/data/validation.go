package data

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator"
)

// ValidationError wraps the validators FieldError so we do not
// expose this to out code
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
	validate.RegisterValidation("sku", ValidateSKU)

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
	err := v.validate.Struct(i)
	if err == nil {
		return nil
	}

	// Só faz o type assertion se err for do tipo validator.ValidationErrors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		// Se não for, retorna um erro genérico
		return ValidationErrors{ValidationError{nil}}
	}

	var returnErrs []ValidationError
	for _, ve := range validationErrors {
		returnErrs = append(returnErrs, ValidationError{ve})
	}

	return returnErrs
}

// validateSKU
func ValidateSKU(fl validator.FieldLevel) bool {
	sku := fl.Field().String()
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(sku)
	hasDash := strings.Contains(sku, "-")
	return hasNumber && !hasDash
}
