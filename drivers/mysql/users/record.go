package users

import (
	"bansosman/bussiness/users"
	"fmt"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	NIK       int
	Name      string
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
}

func toDomain(rec Users) users.Domain {
	fmt.Println(rec)
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
func fromDomain(domain users.Domain) Users {
	return Users{
		NIK:       domain.NIK,
		Name:      domain.Name,
		Password:  domain.Password,
		FotoRumah: domain.FotoRumah,
		FotoDiri:  domain.FotoDiri,
		Alamat:    domain.Alamat,
	}
}
