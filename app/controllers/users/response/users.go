package response

import (
	"bansosman/bussiness/users"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	NIK       int       `json:"nik"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	FotoRumah string    `json:"foto_rumah"`
	FotoDiri  string    `json:"foto_diri"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
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
