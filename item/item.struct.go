package mapitem

import "time"

type Item[T any] struct {
	Value         T
	Expiration    int64
	createTime    time.Time
	deleteHandler func(v T) error
}

func (item *Item[T]) GetValue() T {
	return item.Value
}

func (item *Item[T]) VerifyExpiration(unixNano int64) bool {
	return item.Expiration > 0 && unixNano > item.Expiration
}

func (item *Item[T]) VerifyTimeDuration(d time.Duration) bool {
	t := item.createTime.Add(d)
	return time.Now().After(t)
}

func (item *Item[T]) DeleteHandler() error {
	if item.deleteHandler != nil {
		return item.deleteHandler(item.Value)
	}
	return nil
}
