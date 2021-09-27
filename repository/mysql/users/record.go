package users

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	NIK       int
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
}

func toDomain() {

}
func fromDomain() {

}
