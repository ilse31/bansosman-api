package apbn

import (
	"time"
)

type Domain struct {
	ID         int
	DanaBansos int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Append(apbns *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(apbns *Domain, id int) (*Domain, error)
	Delete(apbns *Domain, id int) (string, error)
}
type Repository interface {
	Insert(apbns *Domain) (*Domain, error)
	FindAll() ([]Domain, error)
	FindByID(id int) (*Domain, error)
	Update(apbns *Domain, id int) (*Domain, error)
	Delete(apbns *Domain, id int) (string, error)
}
