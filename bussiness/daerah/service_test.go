package daerah_test

import (
	"bansosman/bussiness/daerah"
	_mockDaerah "bansosman/bussiness/daerah/mocks"
	"bansosman/bussiness/geolocation"
	_mockApi "bansosman/bussiness/geolocation/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	geoDomain        geolocation.Domain
	MockgeoRepos     _mockApi.Repository
	daerahRepository _mockDaerah.Repository
	daerahDomain     daerah.Domain
	daerahServe      daerah.Service
)

func TestMain(m *testing.M) {
	daerahServe = daerah.NewServe(&daerahRepository, &MockgeoRepos)
	daerahDomain = daerah.Domain{
		ID:        1,
		ApbnId:    2423,
		Apbns:     2353242342,
		Provinsi:  "awyfawj",
		City:      "adwad",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	m.Run()
}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		daerahRepository.On("Insert", mock.AnythingOfType("*daerah.Domain")).Return(&daerahDomain, nil).Once()

		result, err := daerahServe.Append(&daerahDomain)

		assert.Nil(t, err)
		assert.Equal(t, &daerahDomain, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		daerahRepository.On("Insert", mock.AnythingOfType("*daerah.Domain")).Return(&daerahDomain, assert.AnError).Once()

		_, err := daerahServe.Append(&daerahDomain)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		daerahRepository.On("FindAll").Return([]daerah.Domain{daerahDomain}, nil).Once()

		result, err := daerahServe.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		daerahRepository.On("FindAll").Return([]daerah.Domain{}, assert.AnError).Once()

		_, err := daerahServe.FindAll()

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		daerahRepository.On("Delete", mock.AnythingOfType("*daerah.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := daerahServe.Delete(&daerahDomain, daerahDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		daerahRepository.On("Delete", mock.AnythingOfType("*daerah.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", assert.AnError).Once()

		_, err := daerahServe.Delete(&daerahDomain, daerahDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		daerahRepository.On("FindByID", mock.AnythingOfType("int")).Return(&daerahDomain, nil).Once()

		result, err := daerahServe.FindByID(daerahDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &daerahDomain, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		daerahRepository.On("FindByID", mock.AnythingOfType("int")).Return(&daerahDomain, assert.AnError).Once()

		_, err := daerahServe.FindByID(daerahDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		daerahRepository.On("Update", mock.AnythingOfType("*daerah.Domain"), mock.AnythingOfType("int")).Return(&daerahDomain, nil).Once()

		result, err := daerahServe.Update(&daerahDomain, daerahDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, &daerahDomain, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		daerahRepository.On("Update", mock.AnythingOfType("*daerah.Domain"), mock.AnythingOfType("int")).Return(&daerahDomain, assert.AnError).Once()

		_, err := daerahServe.Update(&daerahDomain, daerahDomain.ID)

		assert.NotNil(t, err)
	})
}
