package request

import (
	"bansosman/bussiness/users"
)

type UsersRegist struct {
	NIK       int    `json:"nik"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	FotoRumah string `json:"foto_rumah"`
	FotoDiri  string `json:"foto_diri"`
	Alamat    string `json:"alamat"`
}

type UsersUpdate struct {
	NIK       int    `json:"nik" `
	Name      string `json:"name"`
	Password  string `json:"password" `
	FotoRumah string `json:"foto_rumah" `
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
		Password:  request.Password,
		FotoRumah: request.FotoRumah,
		FotoDiri:  request.FotoDiri,
		Alamat:    request.Alamat,
	}
}
