package apbn

import (
	"bansosman/bussiness/apbn"

	"gorm.io/gorm"
)

type Apbns struct {
	gorm.Model
	DanaBansos int
}

func ToDomain(rec Apbns) apbn.Domain {
	return apbn.Domain{
		ID:         rec.ID,
		DanaBansos: rec.DanaBansos,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func FromDomain(domain apbn.Domain) Apbns {
	return Apbns{
		DanaBansos: domain.DanaBansos,
	}
}

func ToDomainArray(apbns []Apbns) []apbn.Domain {
	var resp []apbn.Domain

	for _, v := range apbns {
		resp = append(resp, ToDomain(v))
	}
	return resp
}
