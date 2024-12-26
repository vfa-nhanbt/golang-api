package middleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/redis/go-redis/v9"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	"gorm.io/gorm"
)

type contextKey string

const UserIDKey contextKey = "user_id"

type rateLimit int

const RateLimiter rateLimit = 10

func FiberMiddleware(app *fiber.App, db *gorm.DB) {
	app.Use(
		cors.New(),
		logger.New(),
		// GormUserContextMiddleware(db),
		rateLimiterMiddleware(int(RateLimiter), 1*time.Second),
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

func rateLimiterMiddleware(limit int, timeout time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := context.Background()
		ip := c.IP()

		key := fmt.Sprintf("rate_limiter:%s", ip)
		val, err := helpers.GetRedisClient().Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			fmt.Printf("Redis error: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		count, _ := strconv.Atoi(val)
		if count >= limit {
			return c.Status(fiber.StatusTooManyRequests).SendString("Too Many Requests")
		}

		pipe := helpers.GetRedisClient().TxPipeline()
		pipe.Incr(ctx, key)
		if err == redis.Nil || count == 0 {
			pipe.Expire(ctx, key, timeout)
		}
		_, err = pipe.Exec(ctx)
		if err != nil {
			fmt.Printf("Redis error: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		return c.Next()
	}
}
