package xzmap

import mapshared "github.com/averyyan/xz-map/shared"

type Option[K comparable, T any] func(m *Map[K, T])

// 设置分片大小
func WithSharedSize[K comparable, T any](size int) Option[K, T] {
	return func(m *Map[K, T]) {
		m.size = size
		m.shards = make([]*mapshared.Shared[K, T], size)
	}
}
