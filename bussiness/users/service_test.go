package users_test

import (
	"bansosman/app/middleware"
	"bansosman/bussiness/users"
	_usersMock "bansosman/bussiness/users/mocks"
	"bansosman/helper/enkrips"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersService    users.Service
	usersRepository _usersMock.Repository
	usersDomain     users.Domain
	hashedPassword  string
)

func TestMain(m *testing.M) {
	usersService = users.NewService(&usersRepository, &middleware.ConfigJwt{})
	hashedPassword, _ = enkrips.Hash("test")
	usersDomain = users.Domain{
		ID:        1,
		NIK:       231313,
		Name:      "Test Users",
		Email:     "test@email.com",
		Password:  hashedPassword,
		Gaji:      324234,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.Run()
}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		usersRepository.On("Insert", mock.AnythingOfType("*users.Domain")).Return(&usersDomain, nil).Once()

		result, err := usersService.Append(&usersDomain)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		usersRepository.On("Insert", mock.AnythingOfType("*users.Domain")).Return(&usersDomain, assert.AnError).Once()

		_, err := usersService.Append(&usersDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		usersRepository.On("FindAll").Return([]users.Domain{usersDomain}, nil).Once()

		result, err := usersService.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		usersRepository.On("FindAll").Return([]users.Domain{}, assert.AnError).Once()

		_, err := usersService.FindAll()

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		usersRepository.On("Delete", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := usersService.Delete(&usersDomain, usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		usersRepository.On("Delete", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", assert.AnError).Once()

		_, err := usersService.Delete(&usersDomain, usersDomain.ID)

		assert.NotNil(t, err)
	})
}
