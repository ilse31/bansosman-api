package response

import (
	"bansosman/bussiness/users"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	NIK       int       `json:"nik"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FotoRumah string    `json:"foto_rumah"`
	Gaji      int       `json:"gaji"`
	Alamat    string    `json:"alamat"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		NIK:       domain.NIK,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		Gaji:      domain.Gaji,
		Alamat:    domain.Alamat,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func NewResponseArray(user []users.Domain) []Users {
	var response []Users

	for _, val := range user {
		response = append(response, FromDomain(val))
	}
	return response
}
