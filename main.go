package main

import (
	initServices "Redis_test/user/init-services"
	"context"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	log.Println("Job begin")
	log.Println("Okay")

	log.Println("Job begin1")
	log.Println("Job begin2")
	log.Println("Job begin3")
	log.Println("Job begin4")

	if err := godotenv.Load(); err != nil {
		log.Printf("Предупреждение: не удалось загрузить .env файл: %v", err)
	}

	services := initServices.NewContainer()

	ctx := context.Background()

	if err := services.UserUC.CreateUser(ctx, "user:1", "John Doe"); err != nil {
		log.Printf("Error setting value: %v", err)
		return
	}

	log.Println("Successfully set user:1")

	value, err := services.UserUC.GetUser(ctx, "user:1")
	if err != nil {
		log.Printf("Error getting value: %v", err)
		return
	}

	log.Printf("Retrieved value: %s", value)
}
