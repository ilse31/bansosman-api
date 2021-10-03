package apbn

type ServiceAPBn struct {
	repository Repository
}

func NewService(repoApbn Repository) Service {
	return &ServiceAPBn{
		repository: repoApbn,
	}
}

func (servApBn *ServiceAPBn) Append(user *Domain) (*Domain, error) {
	result, err := servApBn.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servApBn *ServiceAPBn) FindAll() ([]Domain, error) {
	result, err := servApBn.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servApBn *ServiceAPBn) FindByID(id int) (*Domain, error) {
	result, err := servApBn.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *ServiceAPBn) Update(user *Domain, id int) (*Domain, error) {
	result, err := servApBn.repository.Update(user, id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servApBn *ServiceAPBn) Delete(user *Domain, id int) (string, error) {
	result, err := servApBn.repository.Delete(user, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
