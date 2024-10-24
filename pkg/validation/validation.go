package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func isNotNil(fl validator.FieldLevel) bool {
	fmt.Println(fl.Field())
	return true
}

func RegisterCustomValidations() {
	Validate.RegisterValidation("notnil", isNotNil)
}
