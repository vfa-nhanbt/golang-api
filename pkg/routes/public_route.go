package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vfa-nhanbt/todo-api/app/controllers"
)

// PublicRoutes func for describe group of public routes
func PublicRoutes(a *fiber.App) {
	// Create routes group
	route := a.Group("/api/v1")

	/// Auth Route
	route.Post("/auth/sign-up", controllers.GetAuthController().SignUpHandler)
	route.Post("/auth/sign-in", controllers.GetAuthController().SignInHandler)
	// route.Post("/test/add-followed-author", controllers.GetAuthController().AddFollowedAuthor)

	/// Book Route
	route.Get("/book/get-all", controllers.GetBookController().GetAllBooks)
	route.Get("/book/get/id/:id", controllers.GetBookController().GetBookByID)
	route.Get("/book/get/filter/", controllers.GetBookController().GetBooksByPage)
	route.Get("/book/search/", controllers.GetBookController().SearchBooks)

	/// Mail Route
	route.Post("/mail/send", controllers.GetSendEmailController().SendEmail)

	/// Notification Route
	route.Post("/notification/send", controllers.GetSendNotificationController().SendNotification)
	route.Post("/notification/topic/subscribe", controllers.GetSendNotificationController().SubscribeToTopic)
	route.Post("/notification/topic/unsubscribe", controllers.GetSendNotificationController().UnsubscribeToTopic)
}
