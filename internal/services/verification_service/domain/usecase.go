package domain

import (
	"encoding/json"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"pushpost/internal/services/verification_service/entity"
	"pushpost/pkg/kafka"
	"pushpost/pkg/utils"
)

type VerificationUseCase struct {
	kafkaProducer kafka.Producer
}

func (u *VerificationUseCase) OTPVerificationRequest(msg kafkago.Message) {
	var request entity.VerificationRequest

	err := json.Unmarshal(msg.Value, &request)

	if err != nil {
		log.Println("error parsing json:", err)

		fmt.Println(err)
	}

	newOTP := utils.NewOTP()

	otpMessage := entity.OTPMessage{
		Email:  request.Email,
		OTP:    newOTP.Code,
		Expiry: newOTP.Expiry,
	}

	messageBytes, err := json.Marshal(otpMessage)

	if err != nil {

		log.Println("error serialization otp message:", err)

		fmt.Println(err)
	}

	err = u.kafkaProducer.SendMessage(messageBytes)

	if err != nil {
		log.Println("error sending otp to kafka:", err)
	} else {
		fmt.Printf("otp %s sent to kafka for %s \n", newOTP, request.Email)
	}

	fmt.Println(err)
}
