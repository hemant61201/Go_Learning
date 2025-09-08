package repositary

import (
	"Simple_Programs/Day10-11/MockTest/model"

	"github.com/stretchr/testify/mock"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) FindById(id int64) (*model.Book, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*model.Book), args.Error(1)
}
