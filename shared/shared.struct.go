package xzmapshared

import (
	"sync"

	mapitem "github.com/averyyan/xz-map/item"
)

// Map 分片
type Shared[K comparable, T any] struct {
	sync.RWMutex
	items map[K]mapitem.Item[T]
}

// 分片设置值
func (shared *Shared[K, T]) SetItem(key K, value mapitem.Item[T]) {
	shared.Lock()
	defer shared.Unlock()
	shared.items[key] = value
}

// 分片获取值
func (shared *Shared[K, T]) GetItem(key K) (mapitem.Item[T], bool) {
	shared.RLock()
	defer shared.RUnlock()
	value, ok := shared.items[key]
	return value, ok
}

// 获取分片值
func (shared *Shared[K, T]) GetItems() map[K]mapitem.Item[T] {
	return shared.items
}

// 分片是否存在
func (shared *Shared[K, T]) HasItem(key K) bool {
	shared.RLock()
	defer shared.RUnlock()
	_, ok := shared.items[key]
	return ok
}

// 分片移除值
func (shared *Shared[K, T]) RemoveItem(key K) (K, mapitem.Item[T]) {
	shared.Lock()
	defer shared.Unlock()
	value := shared.items[key]
	value.DeleteHandler() // 删除函数执行
	delete(shared.items, key)
	return key, value
}

// 分片数据量
func (shared *Shared[K, T]) CountItems() int {
	shared.RLock()
	defer shared.RUnlock()
	return len(shared.items)
}
