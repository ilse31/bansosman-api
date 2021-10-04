package covid

type Domain struct {
	FID        int
	Kode_Provi int
	Provinsi   string
	Kasus_Posi int
	Kasus_Semb int
	Kasus_Meni int
}

type Repository interface {
	GetCaseCovid() (Domain, error)
}
