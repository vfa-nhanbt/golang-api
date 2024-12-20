package helpers

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetDBContext(c *fiber.Ctx) (context.Context, error) {
	dbWithCtx, ok := c.Locals("DB").(*gorm.DB)
	if !ok {
		return nil, errors.New("db not found from fiber context")
	}
	ctx := dbWithCtx.Statement.Context
	return ctx, nil
}
