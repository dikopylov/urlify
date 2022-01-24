package service

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/domain/reference/repository"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Insert(entity *model.Reference) error {
	args := m.Called(entity)
	return args.Error(0)
}

func (m *RepositoryMock) GetByCriteria(criteria repository.Criteria) (*model.Reference, error) {
	args := m.Called(criteria)
	return args.Get(0).(*model.Reference), args.Error(1)
}

type FactoryMock struct {
	mock.Mock
}

func (m *FactoryMock) Make(url string) *model.Reference {
	args := m.Called(url)
	return args.Get(0).(*model.Reference)
}

func TestReferenceService_Encode_ReferenceDoesntExists(t *testing.T) {
	repositoryMock := new(RepositoryMock)
	factoryMock := new(FactoryMock)

	link := "https://test.com"
	newReference := &model.Reference{}
	var nilReference *model.Reference
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnUrl, link)

	repositoryMock.On("GetByCriteria", criteria).Return(nilReference, sql.ErrNoRows)
	factoryMock.On("Make", link).Return(newReference, nil)
	repositoryMock.On("Insert", newReference).Return(nil)

	service := NewReferenceService(repositoryMock, factoryMock)
	actual, actualErr := service.Encode(link)
	expected := newReference
	var expectedErr interface{} = nil

	assert.Equal(t, actual, expected)
	assert.Equal(t, actualErr, expectedErr)
	repositoryMock.AssertExpectations(t)
	factoryMock.AssertExpectations(t)
}

func TestReferenceService_Encode_ReferenceDoesntExists_FailInsert(t *testing.T) {
	repositoryMock := new(RepositoryMock)
	factoryMock := new(FactoryMock)

	link := "https://test.com"
	newReference := &model.Reference{}
	var nilReference *model.Reference
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnUrl, link)

	repositoryMock.On("GetByCriteria", criteria).Return(nilReference, sql.ErrNoRows)
	factoryMock.On("Make", link).Return(newReference, nil)
	repositoryMock.On("Insert", newReference).Return(sql.ErrTxDone)

	service := NewReferenceService(repositoryMock, factoryMock)
	actual, actualErr := service.Encode(link)
	expectedErr := sql.ErrTxDone

	assert.Nil(t, actual)
	assert.Equal(t, actualErr, expectedErr)
	repositoryMock.AssertExpectations(t)
	factoryMock.AssertExpectations(t)
}
