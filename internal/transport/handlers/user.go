package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/domain/dto"
	"pushpost/internal/domain/usecase"
	"pushpost/internal/entity"
)

type UserHandler struct {
	useCase usecase.UserUseCase
}

func RegisterUserHandler(useCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{useCase: useCase}
}

func (u *UserHandler) RegisterUser(c *fiber.Ctx) error {

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

	err := u.useCase.RegisterUser(&params)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}
