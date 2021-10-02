package daerah

type ServeDaerah struct {
	repository Repository
}

func NewService(repoDaerah Repository) Service {
	return &ServeDaerah{
		repository: repoDaerah,
	}
}

func (servDaerah *ServeDaerah) Append(user *Domain) (*Domain, error) {
	result, err := servDaerah.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servDaerah *ServeDaerah) FindAll() ([]Domain, error) {
	result, err := servDaerah.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servDaerah *ServeDaerah) FindByID(id int) (*Domain, error) {
	result, err := servDaerah.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servDaerah *ServeDaerah) Update(user *Domain) (*Domain, error) {
	result, err := servDaerah.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servDaerah *ServeDaerah) Delete(user *Domain, id int) (string, error) {
	result, err := servDaerah.repository.Delete(user, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
