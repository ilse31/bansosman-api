package request

import (
	"bansosman/bussiness/daerah"
	"time"
)

type Daerah struct {
	ID        int       `json:"id"`
	IdApbn    int       `json:"apbn_id"`
	Provinsi  string    `json:"provinsi"`
	Kabupaten string    `json:"kabupaten"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(req Daerah) *daerah.Domain {
	return &daerah.Domain{
		ID:        int(req.ID),
		IdApbn:    req.IdApbn,
		Provinsi:  req.Provinsi,
		Kabupaten: req.Kabupaten,
	}
}
