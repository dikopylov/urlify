package service

import (
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
	"urlify/internal/model/domain/reference/factories"
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

func TestReferenceService_Encode(t *testing.T) {
	link := "https://test.com"

	repositoryMock := new(RepositoryMock)
	//repositoryMock.On("Encode", mock.Anything).Return(url)

	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnUrl, link)

	reference, err := repositoryMock.GetByCriteria(criteria)

	factory := new(FactoryMock)

	service := NewReferenceService(repositoryMock, factory)
	service.Encode(link)
}

func TestReferenceService_GetByCriteria(t *testing.T) {
	type fields struct {
		repository repository.ReferenceRepository
		factory    factories.ReferenceFactory
	}
	type args struct {
		criteria repository.Criteria
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Reference
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &ReferenceService{
				repository: tt.fields.repository,
				factory:    tt.fields.factory,
			}
			got, err := service.GetByCriteria(tt.args.criteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCriteria() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCriteria() got = %v, want %v", got, tt.want)
			}
		})
	}
}
