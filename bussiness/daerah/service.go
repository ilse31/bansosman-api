package daerah

import (
	"bansosman/bussiness/geolocation"
)

type serviceDaerah struct {
	repository Repository
	geoRepo    geolocation.Repository
}

func NewServe(repoDaerah Repository, geoRepos geolocation.Repository) Service {
	return &serviceDaerah{
		repository: repoDaerah,
		geoRepo:    geoRepos,
	}
}

func (servDaerah *serviceDaerah) Append(daerahs *Domain) (*Domain, error) {
	result, err := servDaerah.repository.Insert(daerahs)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servDaerah *serviceDaerah) FindAll() ([]Domain, error) {
	result, err := servDaerah.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servDaerah *serviceDaerah) FindByID(id int) (*Domain, error) {
	result, err := servDaerah.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servDaerah *serviceDaerah) Update(daerahs *Domain, id int) (*Domain, error) {
	result, err := servDaerah.repository.Update(daerahs, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servDaerah *serviceDaerah) Delete(daerahs *Domain, id int) (string, error) {
	result, err := servDaerah.repository.Delete(daerahs, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}

func (serv *serviceDaerah) GetByIP() ([]Domain2, error) {
	lok, err := serv.geoRepo.GetLocationByIP()
	if err != nil {
		return []Domain2{}, err
	}

	dataDaerah, error1 := serv.repository.FindByCity(lok.City)
	if error1 != nil {
		return []Domain2{}, error1
	}
	return dataDaerah, err
}
