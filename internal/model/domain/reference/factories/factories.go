package factories

import (
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/infrastructure/encoding"
)

type ReferenceFactory struct {
	encoder encoding.EncoderService
}

func NewReferenceFactory(encoder encoding.EncoderService) ReferenceFactory {
	return ReferenceFactory{encoder: encoder}
}

func (factory ReferenceFactory) Make(url string) *model.Reference {
	return &model.Reference{
		Url:  url,
		Hash: factory.encoder.Encode(url),
	}
}
