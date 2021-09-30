package users

import "time"

//Domain Layer->Acuan utama domain
type Domain struct {
	ID        int
	NIK       int
	Name      string
	Email     string
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
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
	Login(name, password string) (string, error)
}

//Ke Database
type Repository interface {
	Insert(user *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
	GetByName(name string) (Domain, error)
}
