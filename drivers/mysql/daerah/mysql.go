package daerah

import (
	"bansosman/bussiness/daerah"

	"gorm.io/gorm"
)

type repoDaerah struct {
	DBConn *gorm.DB
}

func NewRepoMysql(db *gorm.DB) daerah.Repository {
	return &repoDaerah{
		DBConn: db,
	}
}

func (repo *repoDaerah) Insert(apbns *daerah.Domain) (*daerah.Domain, error) {
	recordApbn := FromDomain(*apbns)

	if err := repo.DBConn.Create(&recordApbn).Error; err != nil {
		return &daerah.Domain{}, err
	}
	record, err := repo.FindByID(int(recordApbn.ID))
	if err != nil {
		return &daerah.Domain{}, err
	}
	return record, nil
}

func (repo *repoDaerah) FindByID(id int) (*daerah.Domain, error) {
	var recordApbn Daerahs

	if err := repo.DBConn.Where("apbns.id = ?", id).Find(&recordApbn).Error; err != nil {
		return &daerah.Domain{}, err
	}
	result := ToDomain(recordApbn)
	return &result, nil
}

func (repo *repoDaerah) Update(apbns *daerah.Domain) (*daerah.Domain, error) {
	recordApbn := FromDomain(*apbns)
	if err := repo.DBConn.Where("id=?", recordApbn.ID).Updates(&recordApbn).Error; err != nil {
		return &daerah.Domain{}, err
	}
	record, err := repo.FindByID(int(recordApbn.ID))
	if err != nil {
		return &daerah.Domain{}, err
	}
	return record, err
}
func (repo *repoDaerah) FindAll() ([]daerah.Domain, error) {
	var recordApbn []Daerahs

	if err := repo.DBConn.Find(&recordApbn).Error; err != nil {
		return []daerah.Domain{}, err
	}
	return ToDomainArray(recordApbn), nil
}

func (repo *repoDaerah) Delete(apbns *daerah.Domain, id int) (string, error) {
	recordUsers := FromDomain(*apbns)
	if err := repo.DBConn.Delete(&recordUsers).Error; err != nil {
		return "", err
	}
	return "Deleted.", nil
}
