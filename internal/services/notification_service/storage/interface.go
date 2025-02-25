package storage

import (
	"pushpost/internal/services/notification_service/entity"
)

type NotificationRepository interface {
	CreateNotification(notification *entity.Notification) error
}
