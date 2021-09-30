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
	Password  string `json:"password"`
	FotoRumah string `json:"foto_rumah"`
	FotoDiri  string `json:"foto_diri"`
	Alamat    string `json:"alamat"`
}
type UsersLogin struct {
	NIK      int    `json:"nik"`
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
func ToDomainUpdate(request UsersUpdate) *users.Domain {
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
