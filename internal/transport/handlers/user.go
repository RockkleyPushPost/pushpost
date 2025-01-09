package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/domain/dto"
	"pushpost/internal/domain/usecase"
	"pushpost/internal/entity"
)

type UserHandler struct {
	useCase *usecase.UserUseCase
}

func RegisterUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: &useCase}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {

	var data entity.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	params := dto.RegisterUserDTO{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Age:      data.Age,
	}

	err := h.useCase.RegisterUser(&params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var loginRequest dto.UserLoginDTO
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}

	token, err := h.useCase.Login(loginRequest)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"token": token,
		"type":  "Bearer",
	})
}
