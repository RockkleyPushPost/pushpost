package transport

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/domain/usecase"
	"pushpost/internal/services/user_service/entity"
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

	if err := params.Validate(); err != nil {

		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.useCase.RegisterUser(&params)

	if err != nil {

		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *UserHandler) GetUserByUUID(c *fiber.Ctx) error {
	var data entity.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.useCase.GetByUUID(data.UUID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}

func (h *UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	var data entity.User
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	user, err := h.useCase.GetByEmail(data.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(user)
}

func (h *UserHandler) GetByToken(c *fiber.Ctx) error {
	userUUID := c.Locals("userUUID").(uuid.UUID)

	user, err := h.useCase.GetByUUID(userUUID)
	userData := dto.UserDataByUUID{Name: user.Name, Age: user.Age}

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(200).JSON(userData)
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

func (h *UserHandler) AddFriend(c *fiber.Ctx) error {
	var friendshipRequest struct {
		userUUID    uuid.UUID
		friendEmail string
	}

	if err := c.BodyParser(&friendshipRequest); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format",
		})
	}

	err := h.useCase.AddFriend(email)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Friendship created successfully"})

}
