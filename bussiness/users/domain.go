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
	Append(game *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(game *Domain) (*Domain, error)
	Delete(game *Domain, id int) (string, error)
}

//Ke Database
type Repository interface {
	Insert(game *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(game *Domain) (*Domain, error)
	Delete(game *Domain, id int) (string, error)
}
