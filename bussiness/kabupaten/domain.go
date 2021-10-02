package kabupaten

import "time"

type Domain struct {
	ID        int
	Kabupaten string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Append(kab *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(kab *Domain) (*Domain, error)
	Delete(kab *Domain, id int) (string, error)
}
type Repository interface {
	Insert(kab *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(kab *Domain) (*Domain, error)
	Delete(kab *Domain, id int) (string, error)
}
