package kabupaten

type Servkabupaten struct {
	repository Repository
}

func NewService(repoKab Repository) Service {
	return &Servkabupaten{
		repository: repoKab,
	}
}

func (servApBn *Servkabupaten) Append(user *Domain) (*Domain, error) {
	result, err := servApBn.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servApBn *Servkabupaten) FindAll() ([]Domain, error) {
	result, err := servApBn.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servApBn *Servkabupaten) FindByID(id int) (*Domain, error) {
	result, err := servApBn.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *Servkabupaten) Update(user *Domain) (*Domain, error) {
	result, err := servApBn.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *Servkabupaten) Delete(user *Domain, id int) (string, error) {
	result, err := servApBn.repository.Delete(user, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
