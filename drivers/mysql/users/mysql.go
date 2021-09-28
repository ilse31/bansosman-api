package users

import (
	"bansosman/bussiness/users"

	"gorm.io/gorm"
)

type repoUsers struct {
	DBConnection *gorm.DB
}

func NewRepoMysql(db *gorm.DB) users.Repository {
	return &repoUsers{
		DBConnection: db,
	}
}

func (repo *repoUsers) Register(user *users.Domain) (*users.Domain, error) {
	recordUsers := fromDomain(*user)
	if err := repo.DBConnection.Create(&recordUsers).Error; err != nil {
		return &users.Domain{}, err
	}
	return &users.Domain{}, nil
}
func (repo *repoUsers) Update(user *users.Domain, id int) (*users.Domain, error) {
	return &users.Domain{}, nil
}
func (repo *repoUsers) FindByID(id int) (*users.Domain, error) {
	return &users.Domain{}, nil
}
func (repo *repoUsers) FindAll(generalSearch string, alamat string) []users.Domain {
	return []users.Domain{}
}
func (repo *repoUsers) Delete(id int) (*users.Domain, error) {
	return &users.Domain{}, nil
}
