package users_test

import (
	"bansosman/app/middleware"
	"bansosman/bussiness/users"
	_usersMock "bansosman/bussiness/users/mocks"
	"bansosman/helper/enkrips"
	_enkripsMock "bansosman/helper/enkrips/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersService    users.Service
	usersRepository _usersMock.Repository
	mockEnkrips     _enkripsMock.Helper
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

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		usersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&usersDomain, nil).Once()

		result, err := usersService.FindByID(usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		usersRepository.On("FindByID", mock.AnythingOfType("int")).Return(&usersDomain, assert.AnError).Once()

		_, err := usersService.FindByID(usersDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		usersRepository.On("GetByName", mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		input := users.Domain{
			Name:     "iniadmin",
			Password: "test",
		}

		resp, err := usersService.Login(input.Name, input.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Invalid Test | Wrong Name", func(t *testing.T) {
		usersRepository.On("GetByName", mock.AnythingOfType("string")).Return(users.Domain{}, assert.AnError).Once()

		input := users.Domain{
			Name:     "iniadmin",
			Password: "test",
		}

		resp, err := usersService.Login(input.Name, input.Password)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
	t.Run("Invalid Test | Wrong Password", func(t *testing.T) {
		usersRepository.On("GetByName", mock.AnythingOfType("string")).Return(usersDomain, nil).Once()
		mockEnkrips.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError)

		input := users.Domain{
			ID:        0,
			NIK:       0,
			Name:      "iniadmin",
			Email:     "",
			Password:  "wrong",
			FotoRumah: "",
			Gaji:      0,
			Alamat:    "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		resp, err := usersService.Login(input.Name, input.Password)

		assert.NotNil(t, resp)
		assert.Empty(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		usersRepository.On("Update", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return(&usersDomain, nil).Once()

		usersDomain.Password, _ = enkrips.Hash(usersDomain.Password)
		result, err := usersService.Update(&usersDomain, usersDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &usersDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		usersRepository.On("Update", mock.AnythingOfType("*users.Domain"), mock.AnythingOfType("int")).Return(&usersDomain, assert.AnError).Once()

		usersDomain.Password, _ = enkrips.Hash(usersDomain.Password)
		_, err := usersService.Update(&usersDomain, usersDomain.ID)

		assert.NotNil(t, err)
	})
}
