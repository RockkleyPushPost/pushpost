package transport

import (
	"github.com/gofiber/fiber/v2"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase `bind:"*usecase.AuthUseCase"`
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	var body entity.User

	if err := c.BodyParser(&body); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	params := dto.RegisterUserDTO{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Age:      body.Age,
	}

	if err := params.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.AuthUseCase.RegisterUser(&params)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	var loginRequest dto.UserLoginDTO

	if err := c.BodyParser(&loginRequest); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format:" + err.Error(),
		})
	}

	if err := loginRequest.Validate(); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	verified, err := h.AuthUseCase.IsEmailVerified(loginRequest.Email)

	if err != nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !verified {

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": "verify",
		})
	}

	token, err := h.AuthUseCase.Login(loginRequest)

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

func (h *AuthHandler) VerifyEmailOTP(c *fiber.Ctx) error {
	var body struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	if err := c.BodyParser(&body); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request format:" + err.Error(),
		})
	}

	isValidOTP, err := h.AuthUseCase.VerifyEmailOTP(body.OTP, body.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !isValidOTP {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid otp",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "otp is valid",
	})

}
