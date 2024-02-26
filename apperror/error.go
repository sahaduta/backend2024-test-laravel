package apperror

import "errors"

var (
	// 400
	ErrInvalidInput = errors.New("invalid input")

	// 401
	ErrInvalidUserId = errors.New("invalid user article id")

	// 404
	ErrUserIdNotFound = errors.New("user id not found")

	// 500
	ErrInternal = errors.New("internal error")
)
