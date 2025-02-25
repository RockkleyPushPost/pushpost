package usecase

import (
	"pushpost/internal/services/notification_service/domain"
	"pushpost/internal/services/notification_service/domain/dto"
	"pushpost/internal/services/notification_service/entity"
	"pushpost/internal/services/notification_service/storage"
)

// implementation check
var _ domain.NotificationUseCase = &NotificationUseCase{}

type NotificationUseCase struct {
	NotificationRepo storage.NotificationRepository `bind:"storage.NotificationRepository"`
}

func NewNotificationUseCase(NotificationRepo storage.NotificationRepository) *NotificationUseCase {
	return &NotificationUseCase{NotificationRepo: NotificationRepo}
}

func (u *NotificationUseCase) CreateNotification(dto *dto.CreateNotificationDto) error {
	notification := entity.NewNotification(dto)

	return u.NotificationRepo.CreateNotification(notification)
}
