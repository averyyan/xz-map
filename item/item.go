package mapitem

import "time"

type Item[T any] interface {
	GetValue() T                             // 获取缓存值
	VerifyExpiration(unixNano int64) bool    // 判断是否过期
	VerifyTimeDuration(d time.Duration) bool // 判断是否超过时间
	DeleteHandler() error                    // 删除执行函数
}

// 默认不删除缓存
func New[T any](value T, opts ...Option[T]) *item[T] {
	item := &item[T]{
		value:         value,
		expiration:    0,
		deleteHandler: nil,
		createTime:    time.Now(),
	}
	for _, opt := range opts {
		opt(item)
	}
	return item
}
