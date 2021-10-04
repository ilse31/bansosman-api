package response

import (
	"bansosman/bussiness/daerah"
	"time"
)

type Daerah struct {
	ID        int `json:"id"`
	ApbnId    int `json:"apbn_id"`
	Apbn      int
	Provinsi  string    `json:"provinsi"`
	City      string    `json:"City"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDom(domain daerah.Domain) Daerah {
	return Daerah{
		ID:        domain.ID,
		ApbnId:    domain.ApbnId,
		Apbn:      domain.Apbns,
		Provinsi:  domain.Provinsi,
		City:      domain.City,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func NewResponseArray(daerahs []daerah.Domain) []Daerah {
	var resp []Daerah

	for _, v := range daerahs {
		resp = append(resp, FromDom(v))
	}
	return resp
}
