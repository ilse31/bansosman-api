package admin_test

import (
	"bansosman/app/middleware"
	"bansosman/bussiness/admin"
	_adminMock "bansosman/bussiness/admin/mocks"
	"bansosman/helper/enkrips"
	_enkripsMock "bansosman/helper/enkrips/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockAdminRepo  _adminMock.Repository
	mockEnkrips    _enkripsMock.Helper
	adminService   admin.Service
	hashedPassword string
	adminDom       admin.Domain
)

func TestMain(m *testing.M) {
	adminService = admin.NewadminService(&mockAdminRepo, &middleware.ConfigJwt{})
	hashedPassword, _ = enkrips.Hash("test")
	adminDom = admin.Domain{
		Username: "tes123",
		Password: hashedPassword,
		RoleID:   1,
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAdminRepo.On("Register", mock.Anything).Return(adminDom, nil).Once()

		inputUser := admin.Domain{
			Username: "testUser",
			Password: "test",
			RoleID:   1,
		}

		resp, err := adminService.Register(&inputUser)

		assert.Nil(t, err)
		assert.Equal(t, adminDom, resp)
	})
	t.Run("Invalid Test | Duplicate User", func(t *testing.T) {
		mockAdminRepo.On("Register", mock.Anything).Return(admin.Domain{}, assert.AnError).Once()

		inputUser := admin.Domain{
			Username: "testUser",
			Password: "test",
			RoleID:   1,
		}

		resp, err := adminService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, adminDom, resp)
	})
	t.Run("Invalid Test | Internal Error", func(t *testing.T) {
		mockAdminRepo.On("Register", mock.Anything).Return(admin.Domain{}, assert.AnError).Once()

		inputUser := admin.Domain{
			Username: "testUser",
			Password: "test",
			RoleID:   1,
		}

		resp, err := adminService.Register(&inputUser)

		assert.NotNil(t, err)
		assert.NotEqual(t, adminDom, resp)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(adminDom, nil).Once()

		input := admin.Domain{
			Username: "testUser",
			Password: "test",
		}

		resp, err := adminService.Login(input.Username, input.Password)

		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Invalid Test | Wrong Username", func(t *testing.T) {
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(admin.Domain{}, assert.AnError).Once()

		input := admin.Domain{
			Username: "testUser",
			Password: "test",
		}

		resp, err := adminService.Login(input.Username, input.Password)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
	t.Run("Invalid Test | Wrong Password", func(t *testing.T) {
		mockAdminRepo.On("GetByUsername", mock.AnythingOfType("string")).Return(adminDom, nil).Once()
		mockEnkrips.On("ValidateHash", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("", assert.AnError)

		input := admin.Domain{
			Username: "testUser",
			Password: "wrong",
		}

		resp, err := adminService.Login(input.Username, input.Password)

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		mockAdminRepo.On("GetByID", mock.AnythingOfType("uint")).Return(adminDom, nil).Once()

		resp, err := adminService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, adminDom, resp)
	})
	t.Run("Invalid Test | User Not Found", func(t *testing.T) {
		mockAdminRepo.On("GetByID", mock.AnythingOfType("uint")).Return(admin.Domain{}, assert.AnError).Once()

		resp, err := adminService.GetByID(2)

		assert.NotNil(t, err)
		assert.NotEqual(t, adminDom, resp)
	})
}
