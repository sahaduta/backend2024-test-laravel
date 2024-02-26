package server

import (
	"github.com/sahaduta/backend2024-test-laravel/handler"
	"github.com/sahaduta/backend2024-test-laravel/repository"
	"github.com/sahaduta/backend2024-test-laravel/usecase"
	"gorm.io/gorm"
)

func GetRouterOpts(db *gorm.DB) RouterOpts {

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	opts := RouterOpts{
		UserHandler: userHandler,
	}

	return opts
}
