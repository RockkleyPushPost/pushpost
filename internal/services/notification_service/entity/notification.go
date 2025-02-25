package entity

import (
	"github.com/google/uuid"
	"pushpost/internal/services/notification_service/domain/dto"
	"time"
)

//type NotificationType string
//
//const (
//	NotificationTypeFriendRequest   NotificationType = "friend_request"
//	NotificationTypeMessageReceived NotificationType = "message_received"
//)

type Notification struct {
	UUID      uuid.UUID `json:"uuid"`
	UserUUID  uuid.UUID `json:"userUUID"`
	Type      string    `json:"type"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"created_at"`
}

func NewNotification(dto *dto.CreateNotificationDto) *Notification {
	return &Notification{
		UUID:      uuid.New(),
		UserUUID:  dto.UserUUID,
		Type:      dto.Type,
		Content:   dto.Content,
		IsRead:    false,
		CreatedAt: time.Now(),
	}
}
