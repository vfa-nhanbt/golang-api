package repositories

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BaseResponse struct {
	Code      string      `json:"code"`
	IsSuccess bool        `json:"isSuccess"`
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
	res := BaseResponse{
		Code:      "-1",
		IsSuccess: false,
		Data:      err.Error(),
	}
	return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
}
