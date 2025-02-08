package utils

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := crand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func TokenExpiry() time.Time {
	return time.Now().Add(24 * time.Hour)
}

func GenerateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func OTPExpiry() time.Time {
	return time.Now().Add(5 * time.Minute)
}
