package data

type Treap interface {
	Get(kv *KV) *KV
	Upsert(kv *KV, pri int) Treap
	Delete(kv *KV) Treap
}
