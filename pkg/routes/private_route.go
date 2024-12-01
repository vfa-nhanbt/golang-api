package routes

import (
	"github.com/gofiber/fiber/v2"
	constant "github.com/vfa-nhanbt/todo-api/pkg/constants"
	middleware "github.com/vfa-nhanbt/todo-api/pkg/middleware"
	"github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

// PrivateRoutes func for describe group of private routes
func PrivateRoutes(a *fiber.App) {
	// Create routes group
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Get("/access-check/admin", middleware.JWTProtected([]string{constant.RoleAdmin}), checkAdmin)
	route.Get("/access-check/user", middleware.JWTProtected([]string{constant.RoleViewer}), checkUser)
	route.Get("/access-check/all", middleware.JWTProtected([]string{constant.RoleAdmin, constant.RoleViewer, constant.RoleAdmin}), checkAll)

}

func checkAdmin(c *fiber.Ctx) error {
	res := repositories.BaseResponse{
		Code:      "s-001",
		IsSuccess: true,
		Data:      "is admin",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func checkUser(c *fiber.Ctx) error {
	res := repositories.BaseResponse{
		Code:      "s-001",
		IsSuccess: true,
		Data:      "is user",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func checkAll(c *fiber.Ctx) error {
	res := repositories.BaseResponse{
		Code:      "s-001",
		IsSuccess: true,
		Data:      "is admin and user",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
