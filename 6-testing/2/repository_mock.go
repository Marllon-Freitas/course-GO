// trabalhando com mocks
package tax

import "github.com/stretchr/testify/mock"

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) Save(amount float64) error {
	args := m.Called(amount)
	return args.Error(0)
}
