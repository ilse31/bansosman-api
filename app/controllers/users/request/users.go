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
	Gaji      int    `json:"gaji"`
	Alamat    string `json:"alamat"`
}

type UsersUpdate struct {
	ID        int    `json:"id"`
	NIK       int    `json:"nik"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FotoRumah string `json:"foto_rumah"`
	Gaji      int    `json:"gaji"`
	Alamat    string `json:"alamat"`
}
type UsersLogin struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func ToDomain(request UsersRegist) *users.Domain {
	return &users.Domain{
		NIK:      request.NIK,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Gaji:     request.Gaji,
		Alamat:   request.Alamat,
	}
}
func ToDomainUpdate(request UsersUpdate) *users.Domain {
	return &users.Domain{
		ID:       request.ID,
		NIK:      request.NIK,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Gaji:     request.Gaji,
		Alamat:   request.Alamat,
	}
}
