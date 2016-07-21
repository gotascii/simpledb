package kvdb

import (
	"math/rand"

	"github.com/gotascii/simpledb/aggregate"
	"github.com/gotascii/simpledb/data"
)

type Transaction struct {
	DB     data.Treap
	Counts aggregate.DB
}

func (t *Transaction) Copy() *Transaction {
	return &Transaction{
		DB:     t.DB,
		Counts: t.Counts.Copy(),
	}
}

func (t *Transaction) Get(k string) interface{} {
	p := t.DB.Get(&data.KV{K: k})
	if p != nil {
		return p.V
	}
	return nil
}

func (t *Transaction) Set(k string, v interface{}) {
	kv := t.DB.Get(&data.KV{K: k})

	if kv == nil || kv.V != v {
		t.DB = t.DB.Upsert(&data.KV{k, v}, rand.Int())
		t.Counts.Append(v.(string), 1)
	}

	if kv != nil && kv.V != v {
		t.Counts.Remove(kv.V.(string), 1)
	}
}

func (t *Transaction) Unset(k string) {
	kv := t.DB.Get(&data.KV{K: k})

	if kv != nil {
		t.DB = t.DB.Delete(kv)
		t.Counts.Remove(kv.V.(string), 1)
	}
}

func (t *Transaction) Numequalto(k string) int {
	return t.Counts.Compute(k)
}
