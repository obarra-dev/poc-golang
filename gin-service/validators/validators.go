package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

// ValidateCoolTitle returns true when the field value contains the word "cool".
func ValidateCoolTitle(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Cool")
}
