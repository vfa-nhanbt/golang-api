package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func StartServer(a *fiber.App) {
	/// Get port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		fmt.Print("Cannot get port from environment variable")
	}
	// Run server
	if err := a.Listen(":" + port); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
