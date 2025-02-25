package repository

import (
	"gorm.io/gorm"
	"pushpost/internal/services/notification_service/entity"
)

type NotificationRepository struct {
	DB *gorm.DB `bind:"*gorm.DB"`
}

func NewNotificationRepository(DB *gorm.DB) *NotificationRepository {
	return &NotificationRepository{DB: DB}
}

func (r *NotificationRepository) CreateNotification(notification *entity.Notification) error {
	return r.DB.Create(notification).Error
}
