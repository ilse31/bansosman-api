package daerah

import "time"

type Domain struct {
	ID        int
	ApbnId    int
	Apbns     int
	Provinsi  string
	Kabupaten string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Append(daerahs *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerahs *Domain) (*Domain, error)
	Delete(daerahs *Domain, id int) (string, error)
}
type Repository interface {
	Insert(daerahs *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerahs *Domain) (*Domain, error)
	Delete(daerahs *Domain, id int) (string, error)
}
