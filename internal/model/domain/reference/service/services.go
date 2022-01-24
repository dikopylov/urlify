package service

import (
	"database/sql"
	"urlify/internal/model/domain/reference/factories"
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/domain/reference/repository"
	"urlify/internal/model/domain/reference/rules"
	"urlify/internal/model/infrastructure/validation"
)

type ReferenceService struct {
	repository repository.ReferenceRepository
	factory    factories.ReferenceMaker
}

func NewReferenceService(repository repository.ReferenceRepository, factory factories.ReferenceMaker) ReferenceService {
	return ReferenceService{
		repository: repository,
		factory:    factory,
	}
}

func (service *ReferenceService) Encode(link string) (*model.Reference, error) {
	err := (&validation.Validator{}).
		SetRules(
			[]validation.Rule{
				&validation.LinkIsCorrect{},
				&rules.LinkIsNotEmpty{},
			},
		).
		Validate(link)

	if err != nil {
		return nil, err
	}

	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnUrl, link)

	reference, err := service.repository.GetByCriteria(criteria)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if reference != nil {
		return reference, nil
	}

	reference = service.factory.Make(link)

	err = service.repository.Insert(reference)

	if err != nil {
		return nil, err
	}

	return reference, nil
}

func (service *ReferenceService) GetByCriteria(criteria repository.Criteria) (*model.Reference, error) {
	return service.repository.GetByCriteria(criteria)
}
