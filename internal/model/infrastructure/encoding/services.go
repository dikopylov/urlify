package encoding

import (
	"encoding/base64"
)

type Encoder interface {
	Encode(url string) string
}

type Base64Encoder struct{}

func NewBase64EncoderService() *Base64Encoder {
	return &Base64Encoder{}
}

func (service *Base64Encoder) Encode(url string) string {
	return base64.URLEncoding.EncodeToString([]byte(url))
}
