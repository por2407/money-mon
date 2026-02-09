package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/moneymon/internal/core/domain"
	"github.com/moneymon/internal/core/ports"
)

type AuthHandler struct {
	userService ports.UserService
}

func NewAuthHandler(userService ports.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req domain.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	user, err := h.userService.Register(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Register failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Register success",
		"user":    user,
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req domain.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	token, err := h.userService.Login(c.UserContext(), &req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Login failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})
}
