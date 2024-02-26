package usecase

import (
	"context"
	"errors"

	"github.com/sahaduta/backend2024-test-laravel/apperror"
	"github.com/sahaduta/backend2024-test-laravel/dto"
	"github.com/sahaduta/backend2024-test-laravel/entity"
	"github.com/sahaduta/backend2024-test-laravel/repository"
	"gorm.io/gorm"
)

type UserUsecase interface {
	GetAllUsers(ctx context.Context, payload *dto.UsersRequest) (*dto.UsersResponse, error)
	CreateUser(ctx context.Context, user *entity.User) (uint, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository}
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func (uc *userUsecase) GetAllUsers(ctx context.Context, payload *dto.UsersRequest) (*dto.UsersResponse, error) {
	totalItem, err := uc.userRepository.Count(ctx, payload)
	if err != nil {
		return nil, err
	}

	users, err := uc.userRepository.FindAllUsers(ctx, payload)
	if err != nil {
		return nil, err
	}

	items := make([]*dto.UserResponse, 0)
	for _, v := range users {
		items = append(items, dto.UserToUserResponse(v))
	}

	totalPage := totalItem / payload.Limit
	if totalItem%payload.Limit != 0 {
		totalPage++
	}

	usersResponse := dto.UsersResponse{
		Items:     items,
		TotalItem: totalItem,
		TotalPage: totalPage,
	}
	if err != nil {
		return nil, err
	}

	return &usersResponse, nil
}

func (uc *userUsecase) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	return uc.userRepository.CreateUser(ctx, user)
}

func (uc *userUsecase) UpdateUser(ctx context.Context, user *entity.User) error {
	_, err := uc.userRepository.FindUserDetail(ctx, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrUserIdNotFound
		}
		return err
	}

	return uc.userRepository.UpdateUser(ctx, user)
}

func (uc *userUsecase) DeleteUser(ctx context.Context, user *entity.User) error {
	_, err := uc.userRepository.FindUserDetail(ctx, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.ErrUserIdNotFound
		}
		return err
	}

	return uc.userRepository.DeleteUser(ctx, user)
}
