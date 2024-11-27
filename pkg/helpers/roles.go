package helpers

import (
	"fmt"

	"github.com/vfa-nhanbt/todo-api/pkg/constants"
)

func ValidateRole(role string) error {
	switch role {
	case constants.RoleAdmin, constants.RoleUser:
		return nil
	default:
		return fmt.Errorf("invalid role %s", role)
	}
}
