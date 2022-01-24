package container

import (
	"github.com/jmoiron/sqlx"
	"sync"
	"urlify/internal/model/application"
	"urlify/internal/model/domain/reference/factories"
	"urlify/internal/model/domain/reference/repository"
	reference "urlify/internal/model/domain/reference/service"
	encoding2 "urlify/internal/model/infrastructure/encoding"
)

var (
	once         sync.Once
	appContainer *Container
)

type Container struct {
	db *sqlx.DB
}

func New(database *sqlx.DB) *Container {
	once.Do(func() {
		appContainer = &Container{db: database}
	})

	return appContainer
}

func Get() *Container {
	return appContainer
}

func (container *Container) GetReferenceService() reference.ReferenceService {
	factory := factories.NewReferenceFactory(encoding2.NewBase64EncoderService())

	return reference.NewReferenceService(
		repository.NewPsqlReferenceRepository(container.db),
		&factory,
	)
}

func (container *Container) GetEncoder() application.Encoder {
	return application.NewEncoder(container.GetReferenceService())
}
