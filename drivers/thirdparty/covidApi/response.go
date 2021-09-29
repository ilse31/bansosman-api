package covidApi

<<<<<<< HEAD
=======
import "bansosman/bussiness/covid"

>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
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

<<<<<<< HEAD
// func (resp *Response) toDomain() covid.Domain {
// 	return covid.Domain{
// 		Attributes.FID: resp.Attributes.Fid,
// 	}
// }
=======
func (resp *Response) toDomain() covid.Domain {
	return covid.Domain{
		Attributes.FID: resp.Attributes.Fid,
	}
}
>>>>>>> bb11f13760677954a5180cb5a72ff710b2d46945
