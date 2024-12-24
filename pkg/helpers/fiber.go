package helpers

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func GetContextFromFiber(c *fiber.Ctx) context.Context {
	ctx := context.WithValue(context.Background(), "fiber_ctx", c)
	return ctx
}
