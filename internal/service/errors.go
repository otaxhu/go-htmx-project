package service

import "fmt"

var (
	ErrInternalServer = fmt.Errorf("internal server error")
	ErrNotFound       = fmt.Errorf("resource not found")
	ErrInvalidInput   = fmt.Errorf("the product object passed in is invalid")
)
