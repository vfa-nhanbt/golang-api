package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vfa-nhanbt/todo-api/app/controllers"
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

	/// Route for book api:
	route.Post("/book/create", middleware.JWTProtected([]string{constant.RoleAuthor}), controllers.GetBookController().AddBookHandler)
	route.Get("/book/delete/:id", middleware.JWTProtected([]string{constant.RoleAuthor, constant.RoleAdmin}), controllers.GetBookController().AddBookHandler)
	route.Post("/book/update/:id", middleware.JWTProtected([]string{constant.RoleAuthor}), controllers.GetBookController().UpdateBook)

	/// Route for review api:
	route.Post("/review/create", middleware.JWTProtected([]string{constant.RoleViewer}), controllers.GetReviewController().AddReviewHandler)
	route.Post("/review/delete/:id", middleware.JWTProtected([]string{constant.RoleViewer, constant.RoleAdmin}), controllers.GetReviewController().DeleteReviewByID)
	route.Post("/review/update/:id", middleware.JWTProtected([]string{constant.RoleViewer}), controllers.GetReviewController().UpdateReview)
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
