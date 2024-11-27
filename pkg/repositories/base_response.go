package repositories

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Code      int         `json:"code"`
	IsSuccess bool        `json:"is_success"`
	Data      interface{} `json:"data"`
}

func (r *BaseResponse) ToMap() map[string]interface{} {
	jsonData, err := json.Marshal(r)
	if err != nil {
		log.Printf("Error marshalling %v", err)
		return nil
	}

	var res map[string]interface{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		log.Printf("Error unmarshal %v", err)
		return nil
	}
	return res
}

func BaseErrorResponse(c *fiber.Ctx, err error) error {
	// Return status 400 and error message.
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": true,
		"msg":   err.Error(),
	})
}
