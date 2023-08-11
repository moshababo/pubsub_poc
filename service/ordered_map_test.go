package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	r := require.New(t)

	om := newOrderedMap[string, string]()
	om.Add("key1", "val1")
	om.Add("key2", "val2")
	om.Add("key3", "val3")

	kv := om.Get("key0")
	r.Nil(kv)

	kv = om.Get("key2")
	r.Equal(&KVP[string, string]{"key2", "val2"}, kv)

	om.Remove("key2")
	kv = om.Get("key2")
	r.Nil(kv)

	list := om.GetAll()
	r.Len(list, 2)
	r.Equal(KVP[string, string]{"key1", "val1"}, list[0])
	r.Equal(KVP[string, string]{"key3", "val3"}, list[1])
}
