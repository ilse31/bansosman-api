package response

import (
	"bansosman/bussiness/users"
	"time"
)

type UsersResponse struct {
	ID        int       `json:"id" form:"id"`
	NIK       int       `json:"nik" form:"nik"`
	Name      string    `json:"name" form:"name"`
	Password  string    `json:"password" form:"password"`
	FotoRumah string    `json:"foto_rumah" form:"foto_rumah"`
	FotoDiri  string    `json:"foto_diri" form:"foto_diri"`
	Alamat    string    `json:"alamat" form:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		ID:        domain.ID,
		NIK:       domain.NIK,
		Name:      domain.Name,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
