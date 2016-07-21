package mock

import (
	"github.com/gotascii/simpledb/aggregate"
	extmock "github.com/stretchr/testify/mock"
)

type Counts struct {
	extmock.Mock
}

func (m *Counts) Copy() aggregate.DB {
	m.Called()
	return m
}

func (m *Counts) Append(idx string, val int) {
	m.Called(idx, val)
}

func (m *Counts) Remove(idx string, val int) {
	m.Called(idx, val)
}

func (m *Counts) Compute(val string) int {
	args := m.Called(val)
	return args.Int(0)
}
