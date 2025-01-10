package middleware

import (
	"context"
	"github.com/google/uuid"
)

type contentKey string

const userUUIDKey contentKey = "userUUID"

func WithUserUUID(ctx context.Context, userUUID uuid.UUID) context.Context {
	return context.WithValue(ctx, userUUIDKey, userUUID)
}

func GetUserUUID(ctx context.Context) (uuid.UUID, bool) {
	userUUID, ok := ctx.Value(userUUIDKey).(uuid.UUID)
	return userUUID, ok

}
