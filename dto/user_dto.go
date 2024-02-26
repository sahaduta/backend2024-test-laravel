package dto

import "github.com/sahaduta/backend2024-test-laravel/entity"

type UserRequest struct {
	Name         string `json:"name" binding:"required"`
	Slug         string `json:"slug" binding:"required"`
	IsProject    bool   `json:"is_project" binding:"required"`
	SelfCapture  string `json:"self_capture" binding:"required"`
	ClientPrefix string `json:"client_prefix" binding:"required"`
	ClientLogo   string `json:"client_logo" binding:"required"`
	Address      string `json:"address" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
	City         string `json:"city" binding:"required"`
}

type UsersRequest struct {
	Search string `form:"s"`
	Sort   string `form:"sort"`
	SortBy string `form:"sort-by"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type UserResponse struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsProject    bool   `json:"is_project"`
	SelfCapture  string `json:"self_capture"`
	ClientPrefix string `json:"client_prefix"`
	ClientLogo   string `json:"client_logo"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	City         string `json:"city"`
}

type UsersResponse struct {
	Items     []*UserResponse `json:"items"`
	TotalPage int             `json:"total_page"`
	TotalItem int             `json:"total_item"`
}

func UserToUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		Id:           user.Id,
		Name:         user.Name,
		Slug:         user.Slug,
		IsProject:    user.IsProject,
		SelfCapture:  user.SelfCapture,
		ClientPrefix: user.ClientPrefix,
		ClientLogo:   user.ClientLogo,
		Address:      user.Address,
		PhoneNumber:  user.PhoneNumber,
		City:         user.City,
	}
}
