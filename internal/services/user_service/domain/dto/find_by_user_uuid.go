package dto

import "github.com/google/uuid"

type FindByPairUUID struct {
	FirstUserUUID  uuid.UUID
	SecondUserUUID uuid.UUID
}
