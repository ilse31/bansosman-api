package requset

import "bansosman/bussiness/apbn"

type ApbnReq struct {
	DanaBansos int `json:"dana_bansos"`
}

type ApbnUpd struct {
	ID         int `json:"id"`
	DanaBansos int `json:"dana_bansos"`
}

func ToDomain(req ApbnReq) *apbn.Domain {
	return &apbn.Domain{
		DanaBansos: req.DanaBansos,
	}
}

func ToDomainUpdate(req ApbnUpd) *apbn.Domain {
	return &apbn.Domain{
		ID:         req.ID,
		DanaBansos: req.DanaBansos,
	}
}
