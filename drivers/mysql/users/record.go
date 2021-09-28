package users

import (
	"bansosman/bussiness/users"

	"gorm.io/gorm"
)

type RecordUsers struct {
	gorm.Model
	NIK       int
	Name      string
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
}

func toDomain(rec RecordUsers) users.Domain {
	return users.Domain{
		ID:        int(rec.ID),
		NIK:       rec.NIK,
		Name:      rec.Name,
		Password:  rec.Password,
		FotoRumah: rec.FotoRumah,
		FotoDiri:  rec.FotoDiri,
		Alamat:    rec.Alamat,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}
func fromDomain(domain users.Domain) RecordUsers {
	return RecordUsers{
		NIK:       domain.NIK,
		Name:      domain.Name,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
	}
}
