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

func (servApBn *serviceDaerah) Append(daerahs *Domain) (*Domain, error) {
	result, err := servApBn.repository.Insert(daerahs)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servApBn *serviceDaerah) FindAll() ([]Domain, error) {
	result, err := servApBn.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servApBn *serviceDaerah) FindByID(id int) (*Domain, error) {
	result, err := servApBn.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *serviceDaerah) Update(daerahs *Domain) (*Domain, error) {
	result, err := servApBn.repository.Update(daerahs)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *serviceDaerah) Delete(daerahs *Domain, id int) (string, error) {
	result, err := servApBn.repository.Delete(daerahs, id)
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
