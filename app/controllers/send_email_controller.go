package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/db/repositories"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

type SendEmailController struct {
	Repository *repositories.EmailRepository
}

func (controller *SendEmailController) SendEmail(c *fiber.Ctx) error {
	email := &models.CreateBookEmailModel{}
	/// Validate request body
	err := helpers.ValidateRequestBody(email, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	err = controller.Repository.SendEmail(email)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-send-email-001",
			IsSuccess: false,
			Data:      "Some thing went wrong when sending email!",
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-send-email-001",
		IsSuccess: true,
		Data:      "Email sent successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
