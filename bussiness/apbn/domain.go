package apbn

import (
	"time"
)

type Domain struct {
	ID         uint
	DanaBansos int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Append(daerah *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(daerah *Domain) (*Domain, error)
	Delete(daerah *Domain, id int) (string, error)
}
type Repository interface {
	Insert(user *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	Delete(user *Domain, id int) (string, error)
}
