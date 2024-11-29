package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vfa-nhanbt/todo-api/app/controllers"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

// PublicRoutes func for describe group of public routes
func PublicRoutes(a *fiber.App) {
	// Create routes group
	route := a.Group("/api/v1")

	route.Get("/health-check", checkMongoConnection)
	route.Post("/generate-token", generateTokenForUser)

	/// Auth Route
	route.Post("/auth/sign-up", controllers.GetAuthController().SignUpHandler)
	route.Post("/auth/sign-in", controllers.GetAuthController().SignInHandler)
}

type UserRequestToken struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

func generateTokenForUser(c *fiber.Ctx) error {
	body := &UserRequestToken{}
	// Checking received data from JSON body.
	if err := c.BodyParser(body); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token, err := helpers.GenerateJWT(body.UserID, body.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res := repositories.BaseResponse{
		Code:      "s-001",
		IsSuccess: true,
		Data:      token,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

// / TEST
func checkMongoConnection(ctx *fiber.Ctx) error {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	cluster := os.Getenv("MONGODB_CLUSTER")
	appname := os.Getenv("MONGODB_APPNAME")
	uri := "mongodb+srv://" + url.QueryEscape(username) + ":" +
		url.QueryEscape(password) + "@" + cluster + "/?retryWrites=true&w=majority&appName=" + appname

	docs := "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return err
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	return ctx.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(result),
		"books": jsonData,
	})
}
