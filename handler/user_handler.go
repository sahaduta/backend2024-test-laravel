package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sahaduta/backend2024-test-laravel/apperror"
	"github.com/sahaduta/backend2024-test-laravel/dto"
	"github.com/sahaduta/backend2024-test-laravel/entity"
	"github.com/sahaduta/backend2024-test-laravel/usecase"
)

type UserHandler interface {
	GetAllUsers(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: uc}
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	param := dto.UsersRequest{}
	ctx.ShouldBindQuery(&param)

	sanitizeUsersParam(&param)

	usersResponse, err := h.userUsecase.GetAllUsers(ctx.Request.Context(), &param)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: usersResponse,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	userRequest := dto.UserRequest{}
	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	user := entity.User{
		Name:         userRequest.Name,
		Slug:         userRequest.Slug,
		IsProject:    userRequest.IsProject,
		SelfCapture:  userRequest.SelfCapture,
		ClientPrefix: userRequest.ClientPrefix,
		ClientLogo:   userRequest.ClientLogo,
		Address:      userRequest.Address,
		PhoneNumber:  userRequest.PhoneNumber,
		City:         userRequest.City,
	}
	createdId, err := h.userUsecase.CreateUser(ctx, &user)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp := dto.Response{
		Data: gin.H{"id": createdId},
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	param := ctx.Param("user-id")
	userId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidUserId)
		return
	}

	userRequest := dto.UserRequest{}
	err = ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.Error(apperror.ErrInvalidInput)
		return
	}
	user := entity.User{
		Id:           uint(userId),
		Name:         userRequest.Name,
		Slug:         userRequest.Slug,
		IsProject:    userRequest.IsProject,
		SelfCapture:  userRequest.SelfCapture,
		ClientPrefix: userRequest.ClientPrefix,
		ClientLogo:   userRequest.ClientLogo,
		Address:      userRequest.Address,
		PhoneNumber:  userRequest.PhoneNumber,
		City:         userRequest.City,
	}

	err = h.userUsecase.UpdateUser(ctx, &user)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	param := ctx.Param("user-id")
	userId, err := strconv.Atoi(param)
	if err != nil {
		ctx.Error(apperror.ErrInvalidUserId)
		return
	}

	user := entity.User{Id: uint(userId)}

	err = h.userUsecase.DeleteUser(ctx, &user)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := dto.Response{
		Data: dto.EmptyData{},
	}
	ctx.JSON(http.StatusOK, resp)
}

func sanitizeUsersParam(param *dto.UsersRequest) {
	if param.SortBy != "name" && param.SortBy != "id" {
		param.SortBy = "name"
	}
	if param.Sort != "asc" && param.Sort != "desc" {
		param.Sort = "asc"
	}
	if param.Limit == 0 {
		param.Limit = 10
	}
	if param.Page == 0 {
		param.Page = 1
	}
}
