package users

import (
	"bansosman/bussiness/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	NIK       int
	Name      string
	Email     string
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
}

func ToDomain(rec Users) users.Domain {
	return users.Domain{
		ID:        int(rec.ID),
		NIK:       rec.NIK,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		FotoRumah: rec.FotoRumah,
		FotoDiri:  rec.FotoDiri,
		Alamat:    rec.Alamat,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		NIK:       rec.NIK,
		Name:      rec.Name,
		Email:     rec.Email,
		Password:  rec.Password,
		FotoRumah: rec.FotoRumah,
		FotoDiri:  rec.FotoDiri,
		Alamat:    rec.Alamat,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		NIK:       domain.NIK,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
	}
}
func FromDomainUpdate(domain users.Domain) Users {
	return Users{
		ID:        uint(domain.ID),
		NIK:       domain.NIK,
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
	}
}

func ToDomainArray(user []Users) []users.Domain {
	var response []users.Domain

	for _, val := range user {
		response = append(response, ToDomain(val))
	}
	return response
}
