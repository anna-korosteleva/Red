package useCase

import (
	"Redis_test/user/boundary/repository"
	"context"
)

type UserUseCase struct {
	userRepo repoRedInterface.UserRepository
}

func NewUserUseCase(userRepo repoRedInterface.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, id, name string) error {
	return uc.userRepo.Set(ctx, "user:"+id, name)
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (string, error) {
	return uc.userRepo.Get(ctx, "user:"+id)
}
