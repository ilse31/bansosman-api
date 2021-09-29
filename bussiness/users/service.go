package users

type ServiceUsers struct {
	repository Repository
}

func NewServe(repoUser Repository) Service {
	return &ServiceUsers{
		repository: repoUser,
	}
}

//ini add ke db
func (servUser *ServiceUsers) Append(user *Domain) (*Domain, error) {
	result, err := servUser.repository.Register(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

//update db
func (servUser *ServiceUsers) Update(user *Domain, id int) (*Domain, error) {
	return &Domain{}, nil
}

//searching db
func (servUser *ServiceUsers) FindByID(id int) (*Domain, error) {
	return &Domain{}, nil
}

//searching all
func (servUser *ServiceUsers) FindAll(generalSearch string) []Domain {
	return []Domain{}
}

//Deleted
func (servUser *ServiceUsers) Deleted(id int) (*Domain, error) {
	return &Domain{}, nil
}
