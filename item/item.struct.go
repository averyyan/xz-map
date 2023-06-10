package mapitem

import "time"

type item[T any] struct {
	value         T
	expiration    int64
	createTime    time.Time
	deleteHandler func(v T) error
}

func (item *item[T]) GetValue() T {
	return item.value
}

func (item *item[T]) VerifyExpiration(unixNano int64) bool {
	return item.expiration > 0 && unixNano > item.expiration
}

func (item *item[T]) VerifyTimeDuration(d time.Duration) bool {
	t := item.createTime.Add(d)
	return time.Now().After(t)
}

func (item *item[T]) DeleteHandler() error {
	if item.deleteHandler != nil {
		return item.deleteHandler(item.value)
	}
	return nil
}
