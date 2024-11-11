package main

import (
	"context"
	"github.com/joho/godotenv"
	"hexagonal-architecture-example/internal/adapters/handlers/infrastructure"
	"hexagonal-architecture-example/internal/core/domains/user"
	"hexagonal-architecture-example/internal/core/domains/user/objects"
	"log"
	"os"

	_ "github.com/lib/pq"
	userRepository "hexagonal-architecture-example/internal/adapters/repositories/postgres/user"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres := infrastructure.Sql{
		Host:     os.Getenv("STORAGE_HOST"),
		Port:     os.Getenv("STORAGE_PORT"),
		User:     os.Getenv("STORAGE_USER"),
		Password: os.Getenv("STORAGE_PASSWORD"),
		DBName:   os.Getenv("STORAGE_DB"),
	}
	db := infrastructure.Init(postgres)
	defer db.Close()
	repo := userRepository.NewUserRepository(db)

	ctx := context.Background()

	passwordA, _ := objects.NewPassword("passwordA")
	passwordB, _ := objects.NewPassword("passwordB")
	passwordC, _ := objects.NewPassword("passwordC")
	fakeUsers := []user.User{
		{
			Email:    objects.Email("userA@test.com"),
			Password: passwordA,
		},
		{
			Email:    objects.Email("userB@test.com"),
			Password: passwordB,
		},
		{
			Email:    objects.Email("userC@test.com"),
			Password: passwordC,
		},
	}
	repo.CreateTable(ctx)
	for _, fakeUser := range fakeUsers {
		repo.Store(ctx, fakeUser)
	}

}
