package xzmapitem

import "github.com/golang-module/carbon/v2"

type Option[T any] func(item *item[T])

// 设置删除函数
func WithDeleteHandler[T any](fn func(v T)) Option[T] {
	return func(item *item[T]) {
		item.deleteHandler = fn
	}
}

// 设置过期时间
func WithExpiration[T any](expiration int64) Option[T] {
	return func(item *item[T]) {
		item.expiration = expiration
	}
}

// 无过期时间
func WithNoExpiration[T any]() Option[T] {
	return func(item *item[T]) {
		item.expiration = 0
	}
}

// 设置过期时间(秒)
func WithDurationSeconds[T any](seconds int) Option[T] {
	return func(item *item[T]) {
		item.expiration = carbon.Now().AddSeconds(seconds).Carbon2Time().UnixNano()
	}
}

// 设置过期时间(分钟)
func WithDurationMinutes[T any](minutes int) Option[T] {
	return func(item *item[T]) {
		item.expiration = carbon.Now().AddMinutes(minutes).Carbon2Time().UnixNano()
	}
}

// 设置过期时间(小时)
func WithDurationHours[T any](hours int) Option[T] {
	return func(item *item[T]) {
		item.expiration = carbon.Now().AddHours(hours).Carbon2Time().UnixNano()
	}
}

// 设置过期时间(天)
func WithDurationDays[T any](days int) Option[T] {
	return func(item *item[T]) {
		item.expiration = carbon.Now().AddDays(days).Carbon2Time().UnixNano()
	}
}

// 设置过期时间(月份)
func WithDurationMonths[T any](months int) Option[T] {
	return func(item *item[T]) {
		item.expiration = carbon.Now().AddMonths(months).Carbon2Time().UnixNano()
	}
}
