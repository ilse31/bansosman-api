package admin

import (
	"bansosman/bussiness/admin"

	"gorm.io/gorm"
)

type mysqlRepoAdmin struct {
	Conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) admin.Repository {
	return &mysqlRepoAdmin{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlRepoAdmin) Register(adMins *admin.Domain) (admin.Domain, error) {
	recUser := fromDomain(*adMins)
	err := mysqlRepo.Conn.Create(&recUser).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return recUser.toDomain(), nil
}

func (mysqlRepo *mysqlRepoAdmin) GetByUsername(username string) (admin.Domain, error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "username = ?", username).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (mysqlRepo *mysqlRepoAdmin) GetByID(id uint) (admin.Domain, error) {
	rec := Admins{}
	err := mysqlRepo.Conn.First(&rec, "id = ?", id).Error
	if err != nil {
		return admin.Domain{}, err
	}
	return rec.toDomain(), nil
}
