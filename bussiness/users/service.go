package users

type ServiceUsers struct {
	repository Repository
}

func NewService(repoGame Repository) Service {
	return &ServiceUsers{
		repository: repoGame,
	}
}

func (servUser *ServiceUsers) Append(game *Domain) (*Domain, error) {
	result, err := servUser.repository.Insert(game)
	if err != nil {
		return &Domain{}, err
	}

	return result, nil
}

func (servUser *ServiceUsers) FindAll() ([]Domain, error) {
	result, err := servUser.repository.FindAll()
	if err != nil {
		return []Domain{}, err
	}
	return result, nil
}

func (servUser *ServiceUsers) FindByID(id int) (*Domain, error) {
	result, err := servUser.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *ServiceUsers) Update(game *Domain) (*Domain, error) {
	result, err := servUser.repository.Update(game)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *ServiceUsers) Delete(game *Domain, id int) (string, error) {
	result, err := servUser.repository.Delete(game, id)
	if err != nil {
		return "Fail to delete.", err
	}
	return result, nil
}

func (servUser *ServiceUsers) Login(name, password string) (string, error) {
	// userDom, err := servUser.repository.GetByName(name)
	// if err != nil {
	// 	return "error", err

	// }
	// if !enkrips.ValidateHash(password, userDom.Password) {
	// 	return "error", nil
	// }
	// token := servUser.jwtAuth.GenerateToken(int(userDom.ID))
	// return "token", err
	return "token", nil
}
