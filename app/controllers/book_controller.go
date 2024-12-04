package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/db/repositories"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"

	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
)

type BookController struct {
	Repository *repositories.BookRepository
}

func (controller *BookController) AddBookHandler(c *fiber.Ctx) error {
	bookModel := &models.BookModel{}

	/// Validate request body
	err := helpers.ValidateRequestBody(bookModel, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Insert book to DB
	err = controller.Repository.InsertBook(bookModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-book-001",
			IsSuccess: false,
			Data:      "Cannot insert user record to table with error: " + err.Error(),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Insert success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-book-001",
		IsSuccess: true,
		Data:      bookModel,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *BookController) DeleteBookWithID(c *fiber.Ctx) error {
	var deleteBookRequest struct {
		BookId string `json:"book_id" validate:"required,uuid"`
	}

	/// Validate request body
	err := helpers.ValidateRequestBody(deleteBookRequest, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Get current user id from token
	userId, err := helpers.GetUserIdFromToken(c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Check if user is author of this book
	canDelete, err := controller.Repository.CheckIsAuthor(userId, deleteBookRequest.BookId)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if !canDelete {
		res := pkgRepo.BaseResponse{
			Code:      "e-book-001",
			IsSuccess: false,
			Data:      "user is not owner of this book",
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Delete book from DB
	err = controller.Repository.DeleteBookWithID(deleteBookRequest.BookId)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Delete success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-book-001",
		IsSuccess: true,
		Data:      "Delete book successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *BookController) UpdateBook(c *fiber.Ctx) error {
	bookModel := &models.BookModel{}

	/// Validate request body
	err := helpers.ValidateRequestBody(bookModel, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Get current user id from token
	userId, err := helpers.GetUserIdFromToken(c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Check if user is author of this book
	canUpdate, err := controller.Repository.CheckIsAuthor(userId, bookModel.ID.String())
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	if !canUpdate {
		res := pkgRepo.BaseResponse{
			Code:      "e-book-001",
			IsSuccess: false,
			Data:      "user is not owner of this book",
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	/// Update book
	err = controller.Repository.UpdateBook(bookModel)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	/// Update success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-book-001",
		IsSuccess: true,
		Data:      bookModel,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *BookController) GetAllBooks(c *fiber.Ctx) error {
	books, err := controller.Repository.GetAllBooks()
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	/// Get all books success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-book-001",
		IsSuccess: true,
		Data:      books,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *BookController) GetBookByID(c *fiber.Ctx) error {
	var getBookRequest struct {
		BookID string `json:"book_id" validate:"required,uuid"`
	}

	err := helpers.ValidateRequestBody(getBookRequest, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	books, err := controller.Repository.GetBookByID(getBookRequest.BookID)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}
	/// Get all books success, return status OK
	res := pkgRepo.BaseResponse{
		Code:      "s-book-001",
		IsSuccess: true,
		Data:      books,
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
