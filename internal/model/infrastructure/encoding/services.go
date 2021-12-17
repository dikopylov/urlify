package encoding

import (
	"encoding/base64"
)

type EncoderService interface {
	Encode(url string) string
}

type Base64EncoderService struct{}

func NewBase64EncoderService() *Base64EncoderService {
	return &Base64EncoderService{}
}

func (service *Base64EncoderService) Encode(url string) string {
	return base64.URLEncoding.EncodeToString([]byte(url))
}
