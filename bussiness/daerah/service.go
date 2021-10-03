package daerah

type serviceDaerah struct {
	repository Repository
}

func NewServe(repoDaerah Repository) Service {
	return &serviceDaerah{
		repository: repoDaerah,
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
