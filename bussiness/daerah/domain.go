package daerah

import "time"

type Domain struct {
	ID        int
	IdApbn    int
	Provinsi  string
	Kabupaten string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Append(daerah *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerah *Domain) (*Domain, error)
	Delete(daerah *Domain, id int) (string, error)
}
type Repository interface {
	Insert(daerah *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerah *Domain) (*Domain, error)
	Delete(daerah *Domain, id int) (string, error)
}
