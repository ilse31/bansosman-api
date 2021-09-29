package request

import (
	"bansosman/bussiness/users"
)

type UsersRegist struct {
	NIK       int    `json:"nik"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FotoRumah string `json:"foto_rumah"`
	FotoDiri  string `json:"foto_diri"`
	Alamat    string `json:"alamat"`
}

type UsersUpdate struct {
	ID        int    `json:"id"`
	NIK       int    `json:"nik"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password" `
	FotoRumah string `json:"foto_rumah" `
	FotoDiri  string `json:"foto_diri"`
	Alamat    string `json:"alamat"`
}
type UsersLogin struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

func ToDomain(request UsersRegist) *users.Domain {
	return &users.Domain{
		NIK:       request.NIK,
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		FotoRumah: request.FotoRumah,
		FotoDiri:  request.FotoDiri,
		Alamat:    request.Alamat,
	}
}
<<<<<<< HEAD
func ToDomainUpdate(request UsersUpdate) *users.Domain {
=======

func ToDomainupd(request UsersUpdate) *users.Domain {
>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
	return &users.Domain{
		ID:        request.ID,
		NIK:       request.NIK,
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		FotoRumah: request.FotoRumah,
		FotoDiri:  request.FotoDiri,
		Alamat:    request.Alamat,
	}
}
