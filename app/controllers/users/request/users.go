package request

import "bansosman/bussiness/users"

type UsersRegist struct {
	NIK       int    `json:"nik" form:"nik"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	FotoRumah string `json:"foto_rumah" form:"foto_rumah"`
	FotoDiri  string `json:"foto_diri" form:"foto_diri"`
	Alamat    string `json:"alamat" form:"alamat"`
}

type UsersUpdate struct {
	NIK       int    `json:"nik" form:"nik"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	FotoRumah string `json:"foto_rumah" form:"foto_rumah"`
	FotoDiri  string `json:"foto_diri" form:"foto_diri"`
	Alamat    string `json:"alamat" form:"alamat"`
}

func ToDomain(req UsersRegist) *users.Domain {
	return &users.Domain{
		NIK:       req.NIK,
		Name:      req.Name,
		Password:  req.Password,
		FotoRumah: req.FotoRumah,
		FotoDiri:  req.FotoDiri,
		Alamat:    req.Alamat,
	}
}
