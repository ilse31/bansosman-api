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
func (repo *repoUsers) Insert(game *users.Domain) (*users.Domain, error) {
	recordUsers := FromDomain(*game)
	if err := repo.DBConnection.Create(&recordUsers).Error; err != nil {
		return &users.Domain{}, err
	}

	record, err := repo.FindByID(int(recordUsers.ID))
	if err != nil {
		return &users.Domain{}, err
	}
	return record, nil
}

func (repo *repoUsers) Update(game *users.Domain) (*users.Domain, error) {
	recordUsers := FromDomainUpdate(*game)
	if err := repo.DBConnection.Where("id=?", recordUsers.ID).Updates(&recordUsers).Error; err != nil {
		return &users.Domain{}, err
	}

	record, err := repo.FindByID(int(recordUsers.ID))
	if err != nil {
		return &users.Domain{}, err
	}
	return record, nil
}

func (repo *repoUsers) FindByID(id int) (*users.Domain, error) {
	var recordUsers Users

	if err := repo.DBConnection.Where("users.id = ?", id).Find(&recordUsers).Error; err != nil {
		return &users.Domain{}, err
	}
	result := ToDomain(recordUsers)
	return &result, nil
}

func (repo *repoUsers) FindAll() ([]users.Domain, error) {
	var recordUsers []Users
	if err := repo.DBConnection.Find(&recordUsers).Error; err != nil {
		return []users.Domain{}, err
	}
	return ToDomainArray(recordUsers), nil
}

func (repo *repoUsers) Delete(users *users.Domain, id int) (string, error) {
	recordUsers := FromDomainUpdate(*users)
	if err := repo.DBConnection.Delete(&recordUsers).Error; err != nil {
		return "", err
	}
	return "Deleted.", nil
}
