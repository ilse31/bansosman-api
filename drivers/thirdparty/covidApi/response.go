package covidApi

import "bansosman/bussiness/covid"

type Response struct {
	Attributes struct {
		Fid       int    `json:"FID"`
		KodeProvi int    `json:"Kode_Provi"`
		Provinsi  string `json:"Provinsi"`
		KasusPosi int    `json:"Kasus_Posi"`
		KasusSemb int    `json:"Kasus_Semb"`
		KasusMeni int    `json:"Kasus_Meni"`
	} `json:"attributes"`
}

func (resp *Response) toDomain() covid.Domain {
	return covid.Domain{
		Attributes.FID: resp.Attributes.Fid,
	}
}
