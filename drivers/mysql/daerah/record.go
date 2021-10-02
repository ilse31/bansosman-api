package daerah

import (
	"bansosman/bussiness/daerah"

	"gorm.io/gorm"
)

type Daerahs struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	IdApbn    int    `json:"apbn_id"`
	Provinsi  string `json:"provinsi"`
	Kabupaten string `json:"kabupaten"`
}

func ToDomain(rec Daerahs) daerah.Domain {
	return daerah.Domain{
		ID:        int(rec.ID),
		IdApbn:    rec.IdApbn,
		Provinsi:  rec.Provinsi,
		Kabupaten: rec.Kabupaten,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

// func (rec *Daerahs) toDomain() daerah.Domain {
// 	return daerah.Domain{
// 		IdApbn:      rec.IdApbn,
// 		Provinsi:    rec.Provinsi,
// 		Kabupaten: rec.Kabupaten,
// 	}
// }

func FromDomain(domain daerah.Domain) Daerahs {
	return Daerahs{
		IdApbn:    domain.IdApbn,
		Provinsi:  domain.Provinsi,
		Kabupaten: domain.Kabupaten,
	}
}

func FromDomainUpdate(domain daerah.Domain) Daerahs {
	return Daerahs{
		ID:        uint(domain.ID),
		IdApbn:    domain.IdApbn,
		Provinsi:  domain.Provinsi,
		Kabupaten: domain.Kabupaten,
	}
}

func ToDomainArray(user []Daerahs) []daerah.Domain {
	var response []daerah.Domain

	for _, val := range user {
		response = append(response, ToDomain(val))
	}
	return response
}
