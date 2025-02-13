package dto

import "github.com/google/uuid"

type FindByUserUUIDDto struct {
	FirstUserUUID  uuid.UUID
	SecondUserUUID uuid.UUID
}
