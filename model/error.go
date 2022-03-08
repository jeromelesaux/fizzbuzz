package model

type ErrorMessage struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
