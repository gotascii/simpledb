package data

import "github.com/stretchr/testify/mock"

type MockTreap struct {
	mock.Mock
}

func (m *MockTreap) Get(kv *KV) *KV {
	args := m.Called(kv)
	result, ok := args.Get(0).(*KV)
	if ok {
		return result
	}
	return nil
}

func (m *MockTreap) Upsert(kv *KV, pri int) Treap {
	args := m.Called(kv, pri)
	return args.Get(0).(Treap)
}

func (m *MockTreap) Delete(kv *KV) Treap {
	args := m.Called(kv)
	return args.Get(0).(Treap)
}
