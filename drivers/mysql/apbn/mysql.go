package apbn

import (
	"bansosman/bussiness/apbn"

	"gorm.io/gorm"
)

type repoApbn struct {
	DBConn *gorm.DB
}

func NewRepoMysql(db *gorm.DB) apbn.Repository {
	return &repoApbn{
		DBConn: db,
	}
}

func (repo *repoApbn) Insert(apbns *apbn.Domain) (*apbn.Domain, error) {
	recordApbn := FromDomain(*apbns)

	if err := repo.DBConn.Create(&recordApbn).Error; err != nil {
		return &apbn.Domain{}, err
	}
	record, err := repo.FindByID(int(recordApbn.ID))
	if err != nil {
		return &apbn.Domain{}, err
	}
	return record, nil
}

func (repo *repoApbn) FindByID(id int) (*apbn.Domain, error) {
	var recordApbn Apbns

	if err := repo.DBConn.Where("apbns.id = ?", id).Find(&recordApbn).Error; err != nil {
		return &apbn.Domain{}, err
	}
	result := ToDomain(recordApbn)
	return &result, nil
}

func (repo *repoApbn) Update(apbns *apbn.Domain, id int) (*apbn.Domain, error) {
	recordApbn := FromDomain(*apbns)
	if err := repo.DBConn.Where("id=?", id).Updates(&recordApbn).Error; err != nil {
		return &apbn.Domain{}, err
	}
	record, err := repo.FindByID(int(id))
	if err != nil {
		return &apbn.Domain{}, err
	}
	return record, err
}
func (repo *repoApbn) FindAll() ([]apbn.Domain, error) {
	var recordApbn []Apbns

	if err := repo.DBConn.Find(&recordApbn).Error; err != nil {
		return []apbn.Domain{}, err
	}
	return ToDomainArray(recordApbn), nil
}

func (repo *repoApbn) Delete(apbns *apbn.Domain, id int) (string, error) {
	recordUsers := FromDomainUpdate(*apbns)
	if err := repo.DBConn.Delete(&recordUsers).Error; err != nil {
		return "", err
	}
	return "Deleted.", nil
}
