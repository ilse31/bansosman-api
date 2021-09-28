package bussiness

import "errors"

var (
	errduplicateData = errors.New("Data sudah ada")
	errInvalidLogin  = errors.New("Username/Password salah")
)
