package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vfa-nhanbt/todo-api/app/models"
	"github.com/vfa-nhanbt/todo-api/pkg/helpers"
	pkgRepo "github.com/vfa-nhanbt/todo-api/pkg/repositories"
	"github.com/vfa-nhanbt/todo-api/services/firebase"
)

type SendNotificationController struct {
	Service *firebase.FirebaseMessagingService
}

func (controller *SendNotificationController) SendNotification(c *fiber.Ctx) error {
	notification := &models.NotificationModel{}
	err := helpers.ValidateRequestBody(notification, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	err = controller.Service.SendNotification(notification)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-send-notification-001",
			IsSuccess: false,
			Data:      fmt.Sprintf("Fail to send notification! %v", err),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-send-notification-001",
		IsSuccess: true,
		Data:      "Notification sent successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *SendNotificationController) SubscribeToTopic(c *fiber.Ctx) error {
	subscribeModel := &models.SubscribeToTopicModel{}
	err := helpers.ValidateRequestBody(subscribeModel, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	err = controller.Service.SubscribeToTopic(subscribeModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-subscribe-topic-001",
			IsSuccess: false,
			Data:      fmt.Sprintf("Fail to subscribe to topic! %v", err),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-subscribe-topic-001",
		IsSuccess: true,
		Data:      "Subscribed to topic successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}

func (controller *SendNotificationController) UnsubscribeToTopic(c *fiber.Ctx) error {
	unsubscribeModel := &models.SubscribeToTopicModel{}
	err := helpers.ValidateRequestBody(unsubscribeModel, c)
	if err != nil {
		return pkgRepo.BaseErrorResponse(c, err)
	}

	err = controller.Service.SubscribeToTopic(unsubscribeModel)
	if err != nil {
		res := pkgRepo.BaseResponse{
			Code:      "e-unsubscribe-topic-001",
			IsSuccess: false,
			Data:      fmt.Sprintf("Fail to unsubscribe to topic! %v", err),
		}
		return c.Status(fiber.StatusBadRequest).JSON(res.ToMap())
	}

	res := pkgRepo.BaseResponse{
		Code:      "s-unsubscribe-topic-001",
		IsSuccess: true,
		Data:      "Unsubscribed to topic successfully",
	}
	return c.Status(fiber.StatusOK).JSON(res.ToMap())
}
