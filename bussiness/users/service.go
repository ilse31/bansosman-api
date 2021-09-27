package users

type ServiceUsers struct {
	repository Repository
}

func NewServe(repoUser Repository) Service {
	return ServiceUsers{
		repository: repoUser,
	}
}

func (servUser *ServiceUsers) Append(user *Domain) (*Domain, error) {
	result, err := servUser.repository.Created(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (servUser *ServiceUsers) Update(user *Domain, id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servUser *ServiceUsers) FindByID(id int) (*Domain, error) {
	return &Domain{}, nil
}
func (servUser *ServiceUsers) FindAll(generalSearch string) []Domain {
	return []Domain{}
}
func (servUser *ServiceUsers) Deleted(id int) (*Domain, error) {
	return &Domain{}, nil
}
