package dto

type Response struct {
	Data    any `json:"data,omitempty"`
	Message any `json:"message,omitempty"`
}

type EmptyData struct {
}
