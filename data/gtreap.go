package data

import (
	"strings"

	"github.com/steveyen/gtreap"
)

type Gtreap struct {
	gtreap *gtreap.Treap
}

func (a *Gtreap) Get(kv *KV) *KV {
	kv, ok := a.gtreap.Get(kv).(*KV)
	if ok {
		return kv
	}
	return nil
}

func (a *Gtreap) Upsert(kv *KV, pri int) Treap {
	return &Gtreap{a.gtreap.Upsert(kv, pri)}
}

func (a *Gtreap) Delete(kv *KV) Treap {
	return &Gtreap{a.gtreap.Delete(kv)}
}

func compare(a, b interface{}) int {
	return strings.Compare(a.(*KV).K, b.(*KV).K)
}

func NewGtreap() Treap {
	return &Gtreap{gtreap.NewTreap(compare)}
}
