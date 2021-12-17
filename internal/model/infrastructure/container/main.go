package container

import (
	"github.com/jmoiron/sqlx"
	"urlify/internal/model/domain/reference/factories"
	"urlify/internal/model/domain/reference/repository"
	"urlify/internal/model/domain/reference/services"
	encoding2 "urlify/internal/model/infrastructure/encoding"
)

type Container struct {
	db *sqlx.DB
}

func NewContainer(database *sqlx.DB) Container {
	return Container{db: database}
}

func (container *Container) MakeReferenceService() encoding.ReferenceService {
	return encoding.NewReferenceService(
		repository.NewPsqlReferenceRepository(container.db),
		factories.NewReferenceFactory(encoding2.NewBase64EncoderService()),
	)
}
