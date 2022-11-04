package utils

type AppError struct {
	Error   error
	Message string
	Code    int
}
