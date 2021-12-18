package application

import (
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/domain/reference/repository"
	"urlify/internal/model/domain/reference/service"
)

type Encoder struct {
	service service.ReferenceService
}

func NewEncoder(service service.ReferenceService) Encoder {
	return Encoder{service: service}
}

func (e *Encoder) Encode(link string) (*model.Reference, error) {
	return e.service.Encode(link)
}

func (e *Encoder) Decode(hash string) (*model.Reference, error) {
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnHash, hash)

	return e.service.GetByCriteria(criteria)
}
