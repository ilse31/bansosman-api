package covid

type Domain struct {
	Attributes struct {
		Fid       int    `json:"FID"`
		KodeProvi int    `json:"Kode_Provi"`
		Provinsi  string `json:"Provinsi"`
		KasusPosi int    `json:"Kasus_Posi"`
		KasusSemb int    `json:"Kasus_Semb"`
		KasusMeni int    `json:"Kasus_Meni"`
	} `json:"attributes"`
}

type Repository interface {
	GetAllCaseProv() (Domain, error)
}
