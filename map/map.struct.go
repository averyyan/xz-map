package xzmap

import (
	mapitem "github.com/averyyan/xz-map/item"
	mapshared "github.com/averyyan/xz-map/shared"
)

type Map[K comparable, T any] struct {
	size     int                       // 缓存Map大小
	shards   []*mapshared.Shared[K, T] // 缓存分片
	sharding func(key K) uint32        // 缓存分片策略
	itemOpts []mapitem.Option[T]       // 默认缓存策略
}

// 判断是否有值
func (m *Map[K, T]) Has(key K) bool {
	return m.getShard(key).HasItem(key)
}

// 获取值
func (m *Map[K, T]) Get(key K) (item mapitem.Item[T], ok bool) {
	if m.getShard(key).HasItem(key) {
		item, ok = m.getShard(key).GetItem(key)
		return item, ok
	}
	return item, false
}

// 设置值
func (m *Map[K, T]) Set(key K, value T, opts ...mapitem.Option[T]) {
	dOpts := m.itemOpts
	dOpts = append(dOpts, opts...)
	item := mapitem.New[T](value, dOpts...)
	m.getShard(key).SetItem(key, item)
}

// 删除值
func (m *Map[K, T]) Remove(key K) {
	shard := m.getShard(key)
	if shard.HasItem(key) {
		shard.RemoveItem(key)
	}
}

// 获取所有Item
func (m *Map[K, T]) Items() map[K]mapitem.Item[T] {
	itemMap := make(map[K]mapitem.Item[T])
	for t := range m.IterBuffered() {
		itemMap[t.Key] = t.Val
	}
	return itemMap
}

func (m *Map[K, T]) Size() int {
	size := 0
	for _, shard := range m.shards {
		size += shard.CountItems()
	}
	return size
}

// 非同步 缓存Map清理
func (m *Map[K, T]) Clean() {
	for item := range m.IterBuffered() {
		m.Remove(item.Key)
	}
}

// 获取分片
func (m *Map[K, T]) getShard(key K) *mapshared.Shared[K, T] {
	return m.shards[uint(m.sharding(key))%uint(m.size)]
}

// 多线程读取所有数据
func (m *Map[K, T]) IterBuffered() <-chan tuple[K, mapitem.Item[T]] {
	chans := snapshot(m)
	total := 0
	for _, c := range chans {
		total += cap(c)
	}
	ch := make(chan tuple[K, mapitem.Item[T]], total)
	go fanIn(chans, ch)
	return ch
}
