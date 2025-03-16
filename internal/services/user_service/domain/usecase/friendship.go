package usecase

import (
	"errors"
	"github.com/google/uuid"
	"pushpost/internal/services/user_service/domain"
	"pushpost/internal/services/user_service/domain/dto"
	"pushpost/internal/services/user_service/entity"
	"pushpost/internal/services/user_service/storage"
	"pushpost/pkg/jwt"
)

//type FriendRequestStatus string
//
//const (
//	StatusPending  FriendRequestStatus = "pending"
//	StatusAccepted FriendRequestStatus = "accepted"
//	StatusRejected FriendRequestStatus = "rejected"
//)

// implementation check
var _ domain.FriendshipUseCase = &FriendshipUseCase{}

type FriendshipUseCase struct {
	FriendshipRepo storage.FriendshipRepository `bind:"friendship_repo"`
	UserRepo       storage.UserRepository       `bind:"user_repo"`
	JwtSecret      string
}

func NewFriendshipUseCase(friendshipRequestRepo storage.FriendshipRepository, jwtSecret string) *FriendshipUseCase {

	return &FriendshipUseCase{FriendshipRepo: friendshipRequestRepo, JwtSecret: jwtSecret}
}

func (uc *FriendshipUseCase) CreateFriendshipRequest(prop dto.CreateFriendRequestDto) error {
	claims, err := jwt.VerifyToken(prop.SenderToken, uc.JwtSecret)

	if err != nil {

		return err
	}

	senderUUID := claims["userUUID"].(string)
	recipient, err := uc.UserRepo.GetUserByEmail(prop.RecipientEmail)

	if err != nil {

		return err
	}

	existsDto := dto.FindByPairUUID{FirstUserUUID: uuid.MustParse(senderUUID), SecondUserUUID: recipient.UUID}
	exists, err := uc.FindRequestByPairUUID(existsDto)

	if err != nil {

		return err
	}

	if exists.UUID != uuid.Nil {

		return errors.New("friend request already exists")
	}

	if uuid.MustParse(senderUUID) == recipient.UUID {
		return errors.New("you can't send a friend request to yourself")
	}

	request := entity.NewFriendshipRequest(uuid.MustParse(senderUUID), recipient.UUID)

	return uc.FriendshipRepo.CreateFriendshipRequest(*request)
}

func (uc *FriendshipUseCase) FindRequestByPairUUID(dto dto.FindByPairUUID) (entity.FriendshipRequest, error) {

	return uc.FriendshipRepo.FindRequestByPairUUID(dto)
}

func (uc *FriendshipUseCase) UpdateFriendshipRequestStatus(requestDto dto.UpdateFriendshipRequestDto) error {

	return uc.FriendshipRepo.UpdateFriendshipRequestStatus(requestDto)

}

func (uc *FriendshipUseCase) DeleteFriendshipRequest(dto dto.DeleteFriendshipRequestDto) error {
	//TODO implement me
	panic("implement me")
}

func (uc *FriendshipUseCase) AcceptFriendshipRequest(requestUUID uuid.UUID) error {
	prop := dto.UpdateFriendshipRequestDto{
		RequestUUID: requestUUID,
		Status:      1,
	}

	err := uc.FriendshipRepo.UpdateFriendshipRequestStatus(prop)

	if err != nil {

		return err
	}

	return uc.CreateFriendship(requestUUID)
}

func (uc *FriendshipUseCase) DeclineFriendshipRequest(requestUUID uuid.UUID) error {
	prop := dto.UpdateFriendshipRequestDto{
		RequestUUID: requestUUID,
		Status:      2,
	}

	return uc.FriendshipRepo.UpdateFriendshipRequestStatus(prop)
}

func (uc *FriendshipUseCase) CreateFriendship(requestUUID uuid.UUID) error {
	var existingFriendship entity.Friendship

	request, err := uc.FriendshipRepo.GetRequestByUUID(requestUUID)
	if err != nil {
		return err
	}

	//Check if friendship already exists in either direction
	if existingFriendship, err = uc.FriendshipRepo.FindFriendshipByPairUUID(
		dto.FindByPairUUID{
			FirstUserUUID: request.SenderUUID, SecondUserUUID: request.RecipientUUID}); err != nil {
		return err
	}

	if existingFriendship.UUID != uuid.Nil {
		return errors.New("friendship already exists")
	}

	friendship := entity.Friendship{
		UUID:       uuid.New(),
		UserUUID:   request.SenderUUID,
		FriendUUID: request.RecipientUUID,
	}

	return uc.FriendshipRepo.CreateFriendship(&friendship)
}

func (uc *FriendshipUseCase) FindFriendshipRequest(dto *dto.FindFriendshipRequestDTO) ([]*entity.FriendshipRequest, error) {
	return uc.FriendshipRepo.FindFriendshipRequest(dto)
}

func (uc *FriendshipUseCase) FindFriendships(dto *dto.FindFriendshipDTO) ([]*entity.Friendship, error) {
	return uc.FriendshipRepo.FindFriendships(dto)
}
