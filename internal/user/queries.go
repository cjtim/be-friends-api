package user

import (
	"github.com/cjtim/be-friends-api/repository"
	"github.com/google/uuid"
)

func GetUserInfoById(userID uuid.UUID) (repository.TagUser, error) {
	return repository.UserRepo.GetUserWithTags(userID)
}
