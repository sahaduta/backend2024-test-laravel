package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sahaduta/backend2024-test-laravel/handler"
)

type RouterOpts struct {
	UserHandler handler.UserHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.ContextWithFallback = true

	user := r.Group("/users")
	user.GET("", opts.UserHandler.GetAllUsers)
	user.POST("", opts.UserHandler.CreateUser)
	user.PUT("", opts.UserHandler.UpdateUser)
	user.DELETE("", opts.UserHandler.DeleteUser)

	return r
}
