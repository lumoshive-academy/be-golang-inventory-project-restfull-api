package service

import (
	"go-25-27/model"

	"github.com/stretchr/testify/mock"
)

type MockServiceAuth struct {
	mock.Mock
}

func (m *MockServiceAuth) Login(email, password string) (*model.User, error) {
	args := m.Called(email, password)
	if args.Get(0) != nil {
		return args.Get(0).(*model.User), args.Error(1)
	}
	return nil, args.Error(1)
}
