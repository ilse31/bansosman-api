package admin

import (
	"bansosman/app/middleware"
	"bansosman/bussiness"
	"bansosman/helper/enkrips"
)

type userService struct {
	userRepository Repository
	jwtAuth        *middleware.ConfigJwt
}

func NewUserService(userRepo Repository, jwtauth *middleware.ConfigJwt) Service {
	return &userService{
		userRepository: userRepo,
		jwtAuth:        jwtauth,
	}
}

func (service *userService) Register(userData *Domain) (Domain, error) {

	hashedPassword, _ := enkrips.Hash(userData.Password)
	userData.Password = string(hashedPassword)
	res, err := service.userRepository.Register(userData)
	if res == (Domain{}) {
		return Domain{}, bussiness.ErrDuplicateData
	}
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (service *userService) Login(username, password string) (string, error) {
	userDomain, err := service.userRepository.GetByUsername(username)
	if err != nil {
		return "", bussiness.ErrInvalidLoginInfo
	}

	if !enkrips.ValidateHash(password, userDomain.Password) {
		return "", bussiness.ErrInvalidLoginInfo
	}

	token := service.jwtAuth.GenerateToken(int(userDomain.ID))
	return token, nil
}

func (service *userService) GetByID(id uint) (Domain, error) {
	userDomain, err := service.userRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	return userDomain, nil
}
