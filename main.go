package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"github.com/vfa-nhanbt/todo-api/app/controllers"
	"github.com/vfa-nhanbt/todo-api/db"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"github.com/vfa-nhanbt/todo-api/pkg/routes"
)

func startSever() (*fiber.App, error) {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	fiberConfig := fiber.Config{
		AppName:           "TODO Api Dev v1.0",
		EnablePrintRoutes: true,
		ReadTimeout:       time.Second * time.Duration(readTimeoutSecondsCount),
	}
	app := fiber.New(fiberConfig)
	return app, nil
}

func main() {
	if err := godotenv.Load(".env.dev"); err != nil {
		log.Panic("No env file found")
	}

	/// Connect to DB
	dbClient, err := db.ConnectToDB()
	if err != nil {
		log.Panic("Failed to connect to DB")
	}
	if dbClient.MongoDB != nil {
		defer dbClient.MongoDB.Disconnect(context.Background())
	}

	/// Define Fiber App
	app, err := startSever()
	if err != nil {
		log.Panicf("Cannot start app: %v", err)
	}

	/// Create Controllers
	controllers.InitControllers(dbClient)

	/// Auto migrate db
	db.PostgresAutoMigrate(dbClient.PostgresGormDB)

	/// Config router
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	/// Start server
	helpers.StartServer(app)
}
