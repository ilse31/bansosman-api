package requset

import "bansosman/bussiness/apbn"

type Apbn struct {
	ID         uint `json:"id"`
	DanaBansos int  `json:"dana_bansos"`
}

func ToDomain(req Apbn) *apbn.Domain {
	return &apbn.Domain{
		ID:         req.ID,
		DanaBansos: req.DanaBansos,
	}
}
