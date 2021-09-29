package apbn

import "time"

type Domain struct {
	Id         int
	DanaBansos int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Insert()
}
