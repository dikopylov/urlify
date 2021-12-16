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

func (service ReferenceService) CreateReference(url string) *model.Reference {
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.COLUMN_URL, url)

	reference := service.repository.GetByCriteria(criteria)

	if reference != nil {
		return reference
	}

	reference = service.factory.Make(url)

	service.repository.Insert(reference)

	return reference
}

func (service ReferenceService) GetByHash(hash string) *model.Reference {
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.COLUMN_HASH, hash)

	return service.repository.GetByCriteria(criteria)
}
