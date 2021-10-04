package covidApi

import "bansosman/bussiness/covid"

type Response struct {
	Fid       int    `json:"FID"`
	KodeProvi int    `json:"Kode_Provi"`
	Provinsi  string `json:"Provinsi"`
	KasusPosi int    `json:"Kasus_Posi"`
	KasusSemb int    `json:"Kasus_Semb"`
	KasusMeni int    `json:"Kasus_Meni"`
}

func (resp *Response) toDomain() covid.Domain {
	return covid.Domain{
		FID:        resp.Fid,
		Kode_Provi: resp.KasusPosi,
		Provinsi:   resp.Provinsi,
		Kasus_Posi: resp.KasusPosi,
		Kasus_Semb: resp.KasusSemb,
		Kasus_Meni: resp.KasusMeni,
	}
}
