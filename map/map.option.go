package xzmap

import (
	mapitem "github.com/averyyan/xz-map/item"
	mapshared "github.com/averyyan/xz-map/shared"
)

type Option[K comparable, T any] func(m *Map[K, T])

// 设置分片大小
func WithSharedSize[K comparable, T any](size int) Option[K, T] {
	return func(m *Map[K, T]) {
		m.size = size
		m.shards = make([]*mapshared.Shared[K, T], size)
	}
}

// 配置默认缓存策略
func WithItemOpts[K comparable, T any](opts ...mapitem.Option[T]) Option[K, T] {
	return func(m *Map[K, T]) {
		m.itemOpts = opts
	}
}
