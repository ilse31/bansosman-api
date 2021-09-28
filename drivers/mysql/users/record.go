package users

import (
	"bansosman/bussiness/users"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	NIK       int
	Name      string
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
}

func toDomain(rec Record) users.Domain {
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
func fromDomain(domain users.Domain) Record {
	return Record{
		NIK:       domain.NIK,
		Name:      domain.Name,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
	}
}
