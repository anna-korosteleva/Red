package initServices

import (
	"Redis_test/infrastructure/initInfra"
	"Redis_test/infrastructure/redisRepo"
	"Redis_test/user/useCase"
	"log"
)

type Container struct {
	UserRepo *redisRepo.RedisUserRepository
	UserUC   *useCase.UserUseCase
}

func NewContainer() *Container {

	// Redis client
	redisClient, err := initInfra.NewRedisClient()
	if err != nil {
		log.Fatalf("Failed to create Redis client: %v", err)
	}

	// Repository layer
	userRepo := redisRepo.NewRedisUserRepository(redisClient)

	// Use Case layer
	userUC := useCase.NewUserUseCase(userRepo)

	return &Container{
		UserRepo: userRepo,
		UserUC:   userUC,
	}
}
