package xzmap

import (
	"fmt"

	mapshared "github.com/averyyan/xz-map/shared"
)

type MapStringer interface {
	comparable
	fmt.Stringer
}

func NewStringer[K MapStringer, T any](opts ...Option[K, T]) *Map[K, T] {
	return create(strfnv32[K], opts...)
}

func New[T any](opts ...Option[string, T]) *Map[string, T] {
	return create(fnv32, opts...)
}

func create[K comparable, T any](sharding func(key K) uint32, opts ...Option[K, T]) *Map[K, T] {
	m := &Map[K, T]{
		size:     32,
		sharding: sharding,
	}
	m.shards = make([]*mapshared.Shared[K, T], m.size)
	for _, opt := range opts {
		opt(m)
	}
	for i := 0; i < m.size; i++ {
		m.shards[i] = mapshared.New[K, T]()
	}
	return m
}
