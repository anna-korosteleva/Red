package useCase

import "context"

type UserUseCase interface {
	CreateUser(ctx context.Context, id, name string) error
	GetUser(ctx context.Context, id string) (string, error)
}
