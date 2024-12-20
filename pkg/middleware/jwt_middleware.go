package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	repository "github.com/vfa-nhanbt/todo-api/pkg/repositories"
	"gorm.io/gorm"
)

func JWTProtected(allowedRoles []string, db ...*gorm.DB) func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},
		ContextKey:   "jwt",
		ErrorHandler: handleJWTError,
		SuccessHandler: func(c *fiber.Ctx) error {
			return handleJWTSuccess(c, allowedRoles, db...)
		},
	}

	return jwtMiddleware.New(config)
}

func handleJWTSuccess(c *fiber.Ctx, allowedRoles []string, db ...*gorm.DB) error {
	/// Get current allowRoles from Claims
	userToken := c.Locals("jwt").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	role, ok := claims["role"].(string)

	if !ok {
		res := repository.BaseResponse{
			Code:      "e-003",
			IsSuccess: false,
			Data:      "cannot found roles from credentials",
		}
		return c.Status(fiber.StatusForbidden).JSON(res.ToMap())
	}

	/// Skip block
	for _, allowedRole := range allowedRoles {
		if strings.EqualFold(allowedRole, role) {
			if len(db) > 0 {
				userID, ok := claims["user_id"].(string)
				if !ok {
					fmt.Print("cannot get user_id from token for audit log")
					return c.Next()
				}
				/// Set context for audit log
				ctx := context.WithValue(db[0].Statement.Context, UserIDKey, userID)
				c.Locals("DB", db[0].Session(&gorm.Session{Context: ctx}))
			}

			return c.Next()
		}
	}

	res := repository.BaseResponse{
		Code:      "e-002",
		IsSuccess: false,
		Data:      "Access denied",
	}
	return c.Status(fiber.StatusForbidden).JSON(res.ToMap())
}

func handleJWTError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		res := &repository.BaseResponse{
			Code:      "e-001",
			IsSuccess: false,
			Data:      err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := &repository.BaseResponse{
		Code:      "e-002",
		IsSuccess: false,
		Data:      "Unauthorized! Error with msg: " + err.Error(),
	}
	return c.Status(fiber.StatusUnauthorized).JSON(res.ToMap())
}
