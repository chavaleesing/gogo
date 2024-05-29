package controllers

import (
	"gogo/modules/entities"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	UserUse entities.UserService
}

func (h *userController) Register(c *fiber.Ctx) error {
	req := new(entities.UserRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.UserUse.Register(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

func NewuserController(r fiber.Router, UserUse entities.UserService) {
	controllers := &userController{
		UserUse: UserUse,
	}
	r.Post("/", controllers.Register)
}
