package mapitem

import "time"

// 默认不删除缓存
func New[T any](value T, opts ...Option[T]) *Item[T] {
	item := &Item[T]{
		Value:         value,
		Expiration:    0,
		deleteHandler: nil,
		createTime:    time.Now(),
	}
	for _, opt := range opts {
		opt(item)
	}
	return item
}
