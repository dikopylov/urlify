package factories

import (
	"urlify/internal/model/domain/reference/model"
	"urlify/internal/model/infrastructure/encoding"
)

type ReferenceMaker interface {
	Make(url string) *model.Reference
}

type ReferenceFactory struct {
	encoder encoding.Encoder
}

func NewReferenceFactory(encoder encoding.Encoder) ReferenceFactory {
	return ReferenceFactory{encoder: encoder}
}

func (factory *ReferenceFactory) Make(url string) *model.Reference {
	return &model.Reference{
		Url:  url,
		Hash: factory.encoder.Encode(url),
	}
}
