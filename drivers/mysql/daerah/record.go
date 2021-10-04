package daerah

import (
	"bansosman/bussiness/daerah"
	"bansosman/drivers/mysql/apbn"

	"gorm.io/gorm"
)

type Daerahs struct {
	gorm.Model
	ID       uint `gorm:"primaryKey"`
	ApbnId   int
	Apbns    apbn.Apbns `gorm:"foreignKey:apbn_id"`
	Provinsi string     `json:"provinsi"`
	City     string     `json:"city"`
}

func ToDomain(rec Daerahs) daerah.Domain {
	return daerah.Domain{
		ID:        int(rec.ID),
		ApbnId:    rec.ApbnId,
		Apbns:     rec.Apbns.DanaBansos,
		Provinsi:  rec.Provinsi,
		City:      rec.City,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func (rec *Daerahs) toDomain() daerah.Domain {
	return daerah.Domain{
		ID:       int(rec.ID),
		ApbnId:   rec.ApbnId,
		Provinsi: rec.Provinsi,
		City:     rec.City,
	}
}

func FromDomain(domain daerah.Domain) Daerahs {
	return Daerahs{
		ApbnId:   domain.ApbnId,
		Provinsi: domain.Provinsi,
		City:     domain.City,
	}
}

func FromDomainUpdate(domain daerah.Domain) Daerahs {
	return Daerahs{
		ID:       uint(domain.ID),
		ApbnId:   domain.ApbnId,
		Provinsi: domain.Provinsi,
		City:     domain.City,
	}
}

func ToDomainArray(user []Daerahs) []daerah.Domain {
	var response []daerah.Domain

	for _, val := range user {
		response = append(response, ToDomain(val))
	}
	return response
}
