package response

import (
	"bansosman/bussiness/apbn"
	"time"
)

type Apbn struct {
	ID         uint      `json:"id"`
	DanaBansos int       `json:"dana_bansos"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDom(domain apbn.Domain) Apbn {
	return Apbn{
		ID:         uint(domain.ID),
		DanaBansos: domain.DanaBansos,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func NewResponseArray(apbns []apbn.Domain) []Apbn {
	var resp []Apbn

	for _, v := range apbns {
		resp = append(resp, FromDom(v))
	}
	return resp
}
