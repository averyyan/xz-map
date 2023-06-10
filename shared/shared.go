package mapshared

import mapcommon "github.com/averyyan/xz-map/common"

func New[K comparable, T any]() *Shared[K, T] {
	shared := &Shared[K, T]{
		items: make(map[K]mapcommon.MapItem[T]),
	}
	return shared
}
