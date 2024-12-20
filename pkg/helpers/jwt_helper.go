package helpers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Struct to store information in JWT
type TokenData struct {
	UserID      uuid.UUID
	Role        string
	ExpiredTime int64
}

// ----- Generate token block -----
func GenerateJWT(userId string, credential string) (string, error) {
	accessToken, err := generateAccessToken(userId, credential)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func jwtKeyFunc() (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}

func generateAccessToken(userId string, credential string) (string, error) {
	/// Get expires hours for access token
	hoursCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_HOURS_COUNT"))

	claims := jwt.MapClaims{}

	claims["user_id"] = userId
	claims["expire_time"] = time.Now().Add(time.Hour * time.Duration(hoursCount)).Unix()
	claims["role"] = credential

	/// Create a new JWT token with claims
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/// Generate token
	if token, err := payload.SignedString([]byte(os.Getenv("JWT_SECRET_KEY"))); err != nil {
		return "", err
	} else {
		return token, nil
	}
}

// ----- Generate block end -----

// ----- Read token block -----
func ParseJWT(c *fiber.Ctx) (*TokenData, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	/// If token is valid, then extract body and get information from token
	if ok && token.Valid {
		// UserID
		userId, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return nil, err
		}

		// Expires time
		expires := int64(claims["exp"].(float64))

		// Roles
		role, ok := claims["role"].(string)
		if !ok {
			return nil, fmt.Errorf("value is not of type []string")
		}

		return &TokenData{
			UserID:      userId,
			ExpiredTime: expires,
			Role:        role,
		}, nil
	}
	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearerToken := c.Get("Authorization")
	bearerList := strings.Split(bearerToken, " ")
	if len(bearerList) < 2 {
		return ""
	}
	return bearerList[1]
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return jwtKeyFunc()
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

// ----- Read block end -----
func GetUserIdFromToken(c *fiber.Ctx) (string, error) {
	userToken, ok := c.Locals("jwt").(*jwt.Token)
	if !ok {
		return "", errors.New("failed to retrieve JWT token from context")
	}
	claims := userToken.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"].(string)

	if !ok {
		return "", fmt.Errorf("cannot get user_id from token")
	}
	return userId, nil
}
