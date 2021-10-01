package bussiness

import "errors"

var (
	ErrDuplicateData          = errors.New("Data already exist")
	ErrInvalidLoginInfo       = errors.New("Username or password is invalid")
	ErrInternalServer         = errors.New("Something went wrong")
	ErrNearestNotFound        = errors.New("No  found within your area")
	ErrNotFound               = errors.New(" not found")
	ErrUnauthorized           = errors.New("User Unauthorized")
	ErrProductNotFound        = errors.New("Product not found")
	ErrInvalidProductCategory = errors.New("Invalid category")
	ErrOrdersNotFound         = errors.New("No order has made")
	ErrWeightExceed           = errors.New("Weight limit exceeded")
	ErrInvalidPayment         = errors.New("Invalid payment method")
)
