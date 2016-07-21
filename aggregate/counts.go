package aggregate

import (
	"math/rand"

	"github.com/gotascii/simpledb/data"
)

type Counts struct {
	DB data.Treap
}

func (c *Counts) Copy() DB {
	return &Counts{DB: c.DB}
}

func (c *Counts) Append(idx string, new int) {
	count := c.DB.Get(&data.KV{K: idx})
	if count == nil {
		c.DB = c.DB.Upsert(&data.KV{idx, new}, rand.Int())
	} else {
		c.DB = c.DB.Upsert(&data.KV{idx, count.V.(int) + new}, rand.Int())
	}
}

func (c *Counts) Remove(idx string, val int) {
	count := c.DB.Get(&data.KV{K: idx})
	if count.V.(int) <= val {
		c.DB = c.DB.Delete(count)
	} else {
		c.DB = c.DB.Upsert(&data.KV{count.K, count.V.(int) - val}, rand.Int())
	}
}

func (c *Counts) Compute(val string) int {
	kv := c.DB.Get(&data.KV{K: val})
	if kv == nil {
		return 0
	}
	return kv.V.(int)
}
