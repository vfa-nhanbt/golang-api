package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func FiberMiddleware(app *fiber.App, db *gorm.DB) {
	app.Use(
		cors.New(),
		logger.New(),
		// GormUserContextMiddleware(db),
	)
}

func GormUserContextMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := helpers.GetUserIdFromToken(c)
		if err != nil {
			fmt.Print("Cannot get userID from token")
			return c.Next()
		}

		ctx := context.WithValue(db.Statement.Context, UserIDKey, userID)
		c.Locals("DB", db.Session(&gorm.Session{Context: ctx}))

		return c.Next()
	}
}
