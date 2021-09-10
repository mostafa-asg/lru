package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	c := New(3)

	_, ok := c.Get("key1")
	assert.False(t, ok)

	c.Set("key1", 100)
	v, ok := c.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 100, v.(int))
	assert.Equal(t, "key1", c.ll.head.key)
	assert.Equal(t, "key1", c.ll.tail.key)

	c.Set("key2", 200)
	v, ok = c.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 200, v.(int))
	assert.Equal(t, "key2", c.ll.head.key)
	assert.Equal(t, "key1", c.ll.tail.key)

	c.Set("key3", 300)
	v, ok = c.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, 300, v.(int))
	assert.Equal(t, "key3", c.ll.head.key)
	assert.Equal(t, "key1", c.ll.tail.key)

	// key4 must evict key1
	// key2 and key3 must be untact
	c.Set("key4", 400)
	v, ok = c.Get("key4")
	assert.True(t, ok)
	assert.Equal(t, 400, v.(int))
	_, ok = c.Get("key1")
	assert.False(t, ok)
	v, ok = c.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 200, v.(int))
	v, ok = c.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, 300, v.(int))
	assert.Equal(t, "key4", c.ll.head.key)
	assert.Equal(t, "key2", c.ll.tail.key)

	// delete most recently used
	c.Delete("key4")
	_, ok = c.Get("key4")
	assert.False(t, ok)
	assert.Equal(t, 2, len(c.m)) //key2, key3
	assert.Equal(t, "key3", c.ll.head.key)
	assert.Equal(t, "key2", c.ll.tail.key)

	c.Set("key5", 500)
	v, ok = c.Get("key5")
	assert.True(t, ok)
	assert.Equal(t, 500, v.(int))
	assert.Equal(t, "key5", c.ll.head.key)
	assert.Equal(t, "key2", c.ll.tail.key)

	c.Delete("key2")
	c.Delete("key5")
	c.Delete("key3")
	assert.Equal(t, 0, len(c.m))
	assert.Nil(t, c.ll.head)
	assert.Nil(t, c.ll.tail)
}
