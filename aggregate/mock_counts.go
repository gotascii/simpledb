package aggregate

import (
	"github.com/stretchr/testify/mock"
)

type MockCounts struct {
	mock.Mock
}

func (m *MockCounts) Copy() DB {
	m.Called()
	return m
}

func (m *MockCounts) Append(idx string, val int) {
	m.Called(idx, val)
}

func (m *MockCounts) Remove(idx string, val int) {
	m.Called(idx, val)
}

func (m *MockCounts) Compute(val string) int {
	args := m.Called(val)
	return args.Int(0)
}
