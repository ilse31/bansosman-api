package users

import "time"

//Domain Layer->Acuan utama domain
type Domain struct {
	ID        int
	NIK       int
	Name      string
	Password  string
	FotoRumah string
	FotoDiri  string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

//Logical Interface
type Service interface {
	Append(user *Domain) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindAll(generalSearch string) []Domain
	Deleted(id int) (*Domain, error)
}

//Ke Database
type Repository interface {
	Register(user *Domain) (*Domain, error)
	Update(user *Domain, id int) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindAll(generalSearch string, alamat string) []Domain
	Delete(id int) (*Domain, error)
}
