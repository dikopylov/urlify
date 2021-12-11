package encoding

import (
	"urlify/internal/domain/reference/factories"
	"urlify/internal/domain/reference/model"
	"urlify/internal/domain/reference/repository"
)

type ReferenceService struct {
	repository repository.ReferenceRepository
	factory    factories.ReferenceFactory
}

func NewReferenceService(repository repository.ReferenceRepository, factory factories.ReferenceFactory) ReferenceService {
	return ReferenceService{
		repository: repository,
		factory:    factory,
	}
}

func (service ReferenceService) createReference(url string) model.Reference {
	reference := service.factory.Make(url)

	service.repository.Insert(&reference)

	return reference
}
