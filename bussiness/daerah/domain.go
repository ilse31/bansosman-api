package daerah

import (
	"time"
)

type Domain struct {
	ID        int
	ApbnId    int
	Apbns     int
	Provinsi  string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Domain2 struct {
	ID        int
	ApbnId    int
	Provinsi  string
	City      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Append(daerahs *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerah *Domain, id int) (*Domain, error)
	Delete(daerahs *Domain, id int) (string, error)
	GetByIP() ([]Domain2, error)
}
type Repository interface {
	Insert(daerahs *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerah *Domain, id int) (*Domain, error)
	Delete(daerahs *Domain, id int) (string, error)
	FindByCity(city string) ([]Domain2, error)
}
