package xzmapshared

import mapitem "github.com/averyyan/xz-map/item"

func New[K comparable, T any]() *Shared[K, T] {
	shared := &Shared[K, T]{
		items: make(map[K]mapitem.Item[T]),
	}
	return shared
}
