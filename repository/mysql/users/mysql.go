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

func (repo *repoUsers) Created(user *users.Domain) (*users.Domain, error) {
	recordBook := fromDomain(*user)
}
func (repo *repoUsers) Update(user *users.Domain, id int) (*users.Domain, error) {

}
func (repo *repoUsers) FindByID(id int) (*users.Domain, error) {

}
func (repo *repoUsers) FindAll(generalSearch string, alamat string) []users.Domain {

}
func (repo *repoUsers) Delete(id int) (*users.Domain, error) {

}
