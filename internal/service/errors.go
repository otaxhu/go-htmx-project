package service

import "fmt"

var (
	ErrInvalidPageParam     = fmt.Errorf("the page param must be an integer greater or equal than 1")
	ErrInternalServer       = fmt.Errorf("internal server error")
	ErrNotFound             = fmt.Errorf("resource not found")
	ErrInvalidProductObject = fmt.Errorf("the product object passed in is invalid")
)
