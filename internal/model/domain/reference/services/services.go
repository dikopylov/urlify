package encoding

import (
	"fmt"
	"log"
	"net/url"
	"urlify/internal/model/domain/reference/factories"
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/domain/reference/repository"
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

func (service *ReferenceService) CreateReference(link string) *model.Reference {
	parsedLink, err := url.Parse(link)

	if err != nil {
		log.Fatalln(err)
	}

	hostname := parsedLink.Hostname()

	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnUrl, hostname)

	reference := service.repository.GetByCriteria(criteria)

	if reference != nil {
		return reference
	}

	reference = service.factory.Make(hostname)

	if reference.Url == "" {
		log.Print("Error: Empty URL")

		return nil
	}

	fmt.Println("test", reference, hostname)

	service.repository.Insert(reference)

	return reference
}

func (service *ReferenceService) GetByHash(hash string) *model.Reference {
	criteria := repository.Criteria{}
	criteria.AddParameter(repository.ColumnHash, hash)

	return service.repository.GetByCriteria(criteria)
}
