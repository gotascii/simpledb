package mock

import (
	"github.com/gotascii/simpledb/data"
	extmock "github.com/stretchr/testify/mock"
)

type Treap struct {
	extmock.Mock
}

func (m *Treap) Get(kv *data.KV) *data.KV {
	args := m.Called(kv)
	result, ok := args.Get(0).(*data.KV)
	if ok {
		return result
	}
	return nil
}

func (m *Treap) Upsert(kv *data.KV, pri int) data.Treap {
	args := m.Called(kv, pri)
	return args.Get(0).(data.Treap)
}

func (m *Treap) Delete(kv *data.KV) data.Treap {
	args := m.Called(kv)
	return args.Get(0).(data.Treap)
}
