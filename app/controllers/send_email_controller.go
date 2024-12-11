package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
	"github.com/vfa-nhanbt/todo-api/service/mail"
)

type SendEmailController struct {
	Service *mail.EmailService
}

func (controller *SendEmailController) SendEmail(c *fiber.Ctx) error {
	email := &models.CreateBookEmailModel{}
	/// Validate request body
	err := helpers.ValidateRequestBody(email, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	createBookModel := mail.CreateBookEmail{CreateBookModel: *email}

	err = controller.Service.SendEmail(createBookModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-send-email-001",
			IsSuccess: false,
			Data:      fmt.Sprintf("Fail to send email! %v", err),
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
