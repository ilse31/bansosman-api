package apbn_test

import (
	"bansosman/bussiness/apbn"
	_apbnMock "bansosman/bussiness/apbn/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	apbnReposit _apbnMock.Repository
	apbnServ    apbn.Service
	apbnDom     apbn.Domain
)

func TestMain(m *testing.M) {
	apbnServ = apbn.NewService(&apbnReposit)
	apbnDom = apbn.Domain{
		ID:         1,
		DanaBansos: 234346345645,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	m.Run()
}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		apbnReposit.On("Insert", mock.AnythingOfType("*apbn.Domain")).Return(&apbnDom, nil).Once()

		result, err := apbnServ.Append(&apbnDom)

		assert.Nil(t, err)
		assert.Equal(t, &apbnDom, result)
	})

	t.Run("Append | InValid", func(t *testing.T) {
		apbnReposit.On("Insert", mock.AnythingOfType("*apbn.Domain")).Return(&apbnDom, assert.AnError).Once()

		_, err := apbnServ.Append(&apbnDom)

		assert.NotNil(t, err)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("Find All | Valid", func(t *testing.T) {
		apbnReposit.On("FindAll").Return([]apbn.Domain{apbnDom}, nil).Once()

		result, err := apbnServ.FindAll()

		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
	})

	t.Run("Find All | InValid", func(t *testing.T) {
		apbnReposit.On("FindAll").Return([]apbn.Domain{}, assert.AnError).Once()

		_, err := apbnServ.FindAll()

		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Find By ID | Valid", func(t *testing.T) {
		apbnReposit.On("FindByID", mock.AnythingOfType("int")).Return(&apbnDom, nil).Once()

		result, err := apbnServ.FindByID(apbnDom.ID)

		assert.Nil(t, err)
		assert.Equal(t, &apbnDom, result)
	})

	t.Run("Find By ID | InValid", func(t *testing.T) {
		apbnReposit.On("FindByID", mock.AnythingOfType("int")).Return(&apbnDom, assert.AnError).Once()

		_, err := apbnServ.FindByID(apbnDom.ID)

		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		apbnReposit.On("Delete", mock.AnythingOfType("*apbn.Domain"), mock.AnythingOfType("int")).Return("Data Deleted.", nil).Once()

		result, err := apbnServ.Delete(&apbnDom, apbnDom.ID)

		assert.Nil(t, err)
		assert.Equal(t, "Data Deleted.", result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		apbnReposit.On("Delete", mock.AnythingOfType("*apbn.Domain"), mock.AnythingOfType("int")).Return("Fail to delete.", assert.AnError).Once()

		_, err := apbnServ.Delete(&apbnDom, apbnDom.ID)

		assert.NotNil(t, err)
	})
}
