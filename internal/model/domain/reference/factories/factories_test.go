package factories

import (
	"github.com/stretchr/testify/mock"
	"testing"
)

type EncoderMock struct {
	mock.Mock
}

func (m *EncoderMock) Encode(url string) string {
	args := m.Called(url)
	return args.String()
}

func TestReferenceFactory_Make_AssertEncoder(t *testing.T) {
	url := "https://test.com"

	encoder := new(EncoderMock)
	encoder.On("Encode", mock.Anything).Return(url)

	factory := NewReferenceFactory(encoder)
	factory.Make(url)

	encoder.AssertExpectations(t)
}
