package domain

import (
	"pushpost/internal/services/notification_service/domain/dto"
)

type NotificationUseCase interface {
	CreateNotification(dto *dto.CreateNotificationDto) (err error)
}
