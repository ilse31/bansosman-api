package requset

import "bansosman/bussiness/apbn"

type ApbnReq struct {
	DanaBansos int `json:"dana_bansos"`
}

func ToDomain(req ApbnReq) *apbn.Domain {
	return &apbn.Domain{
		DanaBansos: req.DanaBansos,
	}
}
