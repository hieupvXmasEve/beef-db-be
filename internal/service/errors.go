package service

import "errors"

var (
	// ErrNotFound is returned when a requested resource is not found
	ErrNotFound = errors.New("resource not found")
	// ErrInvalidInput is returned when the input parameters are invalid
	ErrInvalidInput = errors.New("invalid input")
	// ErrUnauthorized is returned when the user is not authorized to perform the action
	ErrUnauthorized = errors.New("unauthorized")
)
