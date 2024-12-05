package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/vfa-nhanbt/todo-api/app/models"
	dbRepo "github.com/vfa-nhanbt/todo-api/db/repositories"
	"github.com/vfa-nhanbt/todo-api/pkg/constants"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
)

type AuthController struct {
	Repository *dbRepo.UserRepository
}

func (controller *AuthController) SignUpHandler(c *fiber.Ctx) error {
	signUpBody := &models.SignUpModel{}

	/// Validate request body
	err := helpers.ValidateRequestBody(signUpBody, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

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

	/// Insert account to DB
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

	switch userModel.UserRole {
	case constants.RoleAdmin:
		err = controller.Repository.InsertAdmin(&models.AdminModel{
			UserModel: userModel,
			UserID:    userModel.ID,
		})
	case constants.RoleAuthor:
		err = controller.Repository.InsertAuthor(&models.AuthorModel{
			UserModel: userModel,
			UserID:    userModel.ID,
			Books:     []models.BookModel{},
		})
	case constants.RoleViewer:
		uuidAuthor1, err1 := uuid.Parse("2fded70c-947a-4ee4-bbcd-545f42d27027")
		uuidAuthor2, err2 := uuid.Parse("832bded9-ede0-48e0-aff3-c7ed78317c44")
		if err1 != nil || err2 != nil {
			fmt.Printf("parsing uuid err : %v", err)
		}
		err = controller.Repository.InsertViewer(&models.ViewerModel{
			UserModel: userModel,
			UserID:    userModel.ID,
			FollowedAuthors: []models.AuthorModel{
				{
					UserID: uuidAuthor1,
					UserModel: &models.UserModel{
						ID: uuidAuthor1,
					},
				},
				{
					UserID: uuidAuthor2,
					UserModel: &models.UserModel{
						ID: uuidAuthor2,
					},
				},
			},
		})
	default:
		err = fmt.Errorf("unknown user role: %s", userModel.UserRole)
	}
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
	err := helpers.ValidateRequestBody(signInBody, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

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

	/// Load additional information for different type of users
	switch userFromDB.UserRole {
	case constants.RoleAdmin:
		admin, err := controller.Repository.FindAdminByEmail(userFromDB.Email)
		if err != nil {
			fmt.Printf("Cannot load admin followed admins with error: %v", err)
		} else {
			res.Data = map[string]interface{}{"token": token, "admin": admin}
		}
	case constants.RoleAuthor:
		author, err := controller.Repository.FindAuthorByEmail(userFromDB.Email)
		if err != nil {
			fmt.Printf("Cannot load author followed authors with error: %v", err)
		} else {
			res.Data = map[string]interface{}{"token": token, "author": author}
		}
	case constants.RoleViewer:
		viewer, err := controller.Repository.FindViewerByEmail(userFromDB.Email)
		if err != nil {
			fmt.Printf("Cannot load viewer followed authors with error: %v", err)
		} else {
			res.Data = map[string]interface{}{"token": token, "viewer": viewer}
		}
	}

	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
