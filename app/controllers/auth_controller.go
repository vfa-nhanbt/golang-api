package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vfa-nhanbt/todo-api/app/models"
	dbRepo "github.com/vfa-nhanbt/todo-api/db/repositories"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

type AuthController struct {
	Repository *dbRepo.UserRepository
}

func (controller *AuthController) SignUpHandler(c *fiber.Ctx) error {
	signUpBody := &models.SignUpModel{}
	if err := c.BodyParser(signUpBody); err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Validate body
	validate := helpers.NewValidator()
	if err := validate.Struct(signUpBody); err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if err := helpers.ValidateRole(signUpBody.Role); err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Check if account has already created
	userFromDB, err := controller.Repository.FindUserByID(signUpBody.Email)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if userFromDB != nil {
		res := pkgRepo.BaseResponse{
			Code:      fiber.StatusConflict,
			IsSuccess: false,
			Data:      "Account with this email " + signUpBody.Email + " already exists!",
		}
		return c.Status(fiber.StatusConflict).JSON(res.ToMap())
	}

	/// TODO: Insert account to DB

	return nil
}

func (controller *AuthController) SignInHandler(c *fiber.Ctx) error {
	fmt.Print("Sign In handler")
	return nil
}
