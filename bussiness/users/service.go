package users

type ServiceUsers struct {
	repository Repository
}

func NewServe(repoUser Repository) Service {
	return &ServiceUsers{
		repository: repoUser,
	}
}

func (servGame *ServiceUsers) Append(game *Domain) (*Domain, error) {
	result, err := servGame.repository.Insert(game)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servGame *ServiceUsers) FindAll() ([]Domain, error) {
	result, err := servGame.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servGame *ServiceUsers) FindByID(id int) (*Domain, error) {
	result, err := servGame.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servGame *ServiceUsers) Update(game *Domain) (*Domain, error) {
	result, err := servGame.repository.Update(game)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servGame *ServiceUsers) Delete(game *Domain, id int) (string, error) {
	result, err := servGame.repository.Delete(game, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}
