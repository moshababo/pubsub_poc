package service

import "fmt"

type KVP[K comparable, V any] struct {
	Key K
	Val V
}

func (p KVP[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", p.Key, p.Val)
}

type node[K comparable, V any] struct {
	key K
	val V

	next *node[K, V]
	prev *node[K, V]
}

type orderedMap[K comparable, V any] struct {
	head *node[K, V]
	tail *node[K, V]

	kv map[K]*node[K, V]
}

func newOrderedMap[K comparable, V any]() *orderedMap[K, V] {
	return &orderedMap[K, V]{
		kv: make(map[K]*node[K, V]),
	}
}

func (om *orderedMap[K, V]) Add(key K, val V) {
	// If exists, just update the val.
	if _, ok := om.kv[key]; ok {
		om.kv[key].val = val
		return
	}

	n := &node[K, V]{key: key, val: val}

	if om.head == nil {
		om.head = n
		om.tail = n
	} else {
		om.tail.next = n
		n.prev = om.tail
		om.tail = n
	}

	om.kv[key] = n
}

func (om *orderedMap[K, V]) Remove(key K) {
	n, ok := om.kv[key]
	if !ok {
		return
	}

	if n.prev == nil {
		om.head = n.next
	} else {
		n.prev.next = n.next
	}

	if n.next == nil {
		om.tail = n.prev
	} else {
		n.next.prev = n.prev
	}

	delete(om.kv, key)
}

func (om *orderedMap[K, V]) Get(key K) *KVP[K, V] {
	n, ok := om.kv[key]
	if !ok {
		return nil
	}
	return &KVP[K, V]{n.key, n.val}
}

func (om *orderedMap[K, V]) GetAll() []KVP[K, V] {
	list := make([]KVP[K, V], len(om.kv))
	var i int
	for n := om.head; n != nil; n = n.next {
		list[i] = KVP[K, V]{n.key, n.val}
		i++
	}

	return list
}
