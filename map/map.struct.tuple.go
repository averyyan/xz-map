package xzmap

// 数据快照
type tuple[K comparable, T any] struct {
	Key K
	Val T
}
