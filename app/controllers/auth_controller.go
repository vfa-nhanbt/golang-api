package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	/// Validate request body
	helpers.ValidateRequestBody(signUpBody, c)

	/// Validate registered role
	if err := helpers.ValidateRole(signUpBody.Role); err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Check if account has already created
	userFromDB, err := controller.Repository.FindUserByEmail(signUpBody.Email)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if userFromDB != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-signup-001",
			IsSuccess: false,
			Data:      "Account with this email " + signUpBody.Email + " already exists!",
		}
		return c.Status(fiber.StatusConflict).JSON(res.ToMap())
	}

	/// TODO: Insert account to DB
	userModel := &models.UserModel{
		ID:           uuid.New(),
		Email:        signUpBody.Email,
		Name:         signUpBody.UserName,
		UserRole:     signUpBody.Role,
		PasswordHash: signUpBody.Password,
	}
	err = controller.Repository.InsertUser(userModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-signup-002",
			IsSuccess: false,
			Data:      "Cannot insert user record to table with error: " + err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-signup-001",
		IsSuccess: true,
		Data:      fmt.Sprintf("Insert user record successfully! UserID: %s", userModel.ID),
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *AuthController) SignInHandler(c *fiber.Ctx) error {
	signInBody := &models.SignInModel{}

	/// Validate request body
	helpers.ValidateRequestBody(signInBody, c)

	/// Check if account has already created
	userFromDB, err := controller.Repository.FindUserByEmail(signInBody.Email)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if userFromDB == nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-signin-001",
			IsSuccess: false,
			Data:      fmt.Sprintf("Cannot find User with email %s from DB", signInBody.Email),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Compare user password from database and request body password
	equal := helpers.CompareEncodePassword(signInBody.Password, userFromDB.PasswordHash)
	if !equal {
		res := pkgRepo.BaseResponse{
			Code:      "e-signin-002",
			IsSuccess: false,
			Data:      fmt.Sprintf("Wrong password with email %s", userFromDB.Email),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Generate token from sign in
	token, err := helpers.GenerateJWT(userFromDB.ID.String(), userFromDB.UserRole)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-signup-001",
		IsSuccess: true,
		Data:      map[string]interface{}{"token": token, "user": userFromDB},
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
