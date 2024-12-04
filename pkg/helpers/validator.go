package helpers

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/vfa-nhanbt/todo-api/pkg/constants"
)

func newValidator() *validator.Validate {
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})
	/// Register password validation
	validate.RegisterValidation("password", passwordValidator)
	/// Register user role validation
	validate.RegisterValidation("userRole", userRoleValidator)
	/// Register price validation
	validate.RegisterValidation("updatePrice", updatePriceValidator)
	return validate
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}
	return fields
}

func ValidateRequestBody(body interface{}, c *fiber.Ctx) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}
	/// Validate body
	validate := newValidator()
	if err := validate.Struct(body); err != nil {
		return err
	}
	return nil
}

func passwordValidator(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Password has at least 8 characters
	if len(password) < 8 {
		return false
	}

	// At least 1 lowercase letter
	hasLower := false
	// At least 1 uppercase letter
	hasUpper := false
	// At least 1 number
	hasDigit := false
	// At least 1 special character
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsAny(string(char), "!@#$%^&*(),.?\":{}|<>"):
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}

func userRoleValidator(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	allowedRoles := constants.GetRoles()
	for _, allowedRole := range allowedRoles {
		if strings.EqualFold(allowedRole, role) {
			return true
		}
	}
	return false
}

func updatePriceValidator(fl validator.FieldLevel) bool {
	fmt.Print(fl.Field())
	if fl.Field().IsNil() {
		return true
	}
	if price := fl.Field().Int(); price >= 0 {
		return true
	}
	return false
}
