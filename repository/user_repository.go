package repository

import (
	"context"
	"errors"

	"github.com/sahaduta/backend2024-test-laravel/apperror"
	"github.com/sahaduta/backend2024-test-laravel/dto"
	"github.com/sahaduta/backend2024-test-laravel/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUsers(ctx context.Context, payload *dto.UsersRequest) ([]*entity.User, error)
	FindUserDetail(ctx context.Context, user entity.User) (*entity.User, error)
	Count(ctx context.Context, payload *dto.UsersRequest) (int, error)
	CreateUser(ctx context.Context, user *entity.User) (uint, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAllUsers(ctx context.Context, payload *dto.UsersRequest) ([]*entity.User, error) {
	user := entity.User{}
	users := make([]*entity.User, 0)
	q := r.db.WithContext(ctx).Model(&user).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Order(payload.SortBy + " " + payload.Sort).
		Offset((payload.Page - 1) * payload.Limit)
	err := q.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Count(ctx context.Context, payload *dto.UsersRequest) (int, error) {
	var total int64 = 0
	user := entity.User{}
	err := r.db.WithContext(ctx).Model(user).
		Where("content ILIKE ?", "%"+payload.Search+"%").
		Limit(payload.Limit).
		Count(&total).
		Error
	if err != nil {
		return 0, err
	}
	return int(total), nil
}

func (r *userRepository) FindUserDetail(ctx context.Context, user entity.User) (*entity.User, error) {
	q := r.db.WithContext(ctx).Model(&user).
		Where("users.id = ?", user.Id)
	err := q.First(&user).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, apperror.ErrUserIdNotFound
		default:
			return nil, err
		}
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *entity.User) (uint, error) {
	result := r.db.WithContext(ctx).Model(&entity.User{}).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
