package dto

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

const (
	MinNameLength     = 2
	MaxNameLength     = 32
	MinPasswordLength = 6
	EmailRegex        = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
)

type RegisterUserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func (dto *RegisterUserDTO) Validate() error {
	length := utf8.RuneCountInString(dto.Name)
	if length < MinNameLength || length > MaxNameLength {
		return errors.New("invalid name length")
	}

	if dto.Age < 0 {
		return errors.New("invalid age")
	}

	if !isEmailValid(dto.Email) {
		return errors.New("invalid email")
	}

	if isValid, passwordError := validatePassword(dto.Password); !isValid {
		return errors.New(passwordError)
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(EmailRegex)
	return emailRegex.MatchString(e)
}
func validatePassword(password string) (bool, string) {
	if utf8.RuneCountInString(password) < MinPasswordLength {
		return false, fmt.Sprintf("password must be at least %d symbols long", MinPasswordLength)
	}
	passError := make([]string, 0, 3)
	validations := make(map[string]bool)

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)

	validations["uppercase symbol(s)"] = hasUpper
	validations["lowercase symbol(s)"] = hasLower
	validations["digits"] = hasDigit

	for k, v := range validations {
		if !v {
			passError = append(passError, k)
		}
	}

	if len(passError) == 0 {
		return true, ""
	}

	return false, "password must also contain " + strings.Join(passError, ", ")
}
