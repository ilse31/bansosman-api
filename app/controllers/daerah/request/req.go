package request

import (
	"bansosman/bussiness/daerah"
	"time"
)

type Daerah struct {
	ID        int       `json:"id"`
	ApbnId    int       `json:"apbn_id"`
	Provinsi  string    `json:"provinsi"`
	City      string    `json:"city"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToDomain(req Daerah) *daerah.Domain {
	return &daerah.Domain{
		ID:       int(req.ID),
		ApbnId:   req.ApbnId,
		Provinsi: req.Provinsi,
		City:     req.City,
	}
}
