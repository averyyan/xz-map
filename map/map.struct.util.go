package xzmap

import (
	"fmt"
	"sync"

	mapcommon "github.com/averyyan/xz-map/common"
	mapshared "github.com/averyyan/xz-map/shared"
)

func strfnv32[K fmt.Stringer](key K) uint32 {
	return fnv32(key.String())
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	keyLength := len(key)
	for i := 0; i < keyLength; i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// 获取数据Map的快照
func snapshot[K comparable, T any](m *Map[K, T]) (chans []chan tuple[K, mapcommon.MapItem[T]]) {
	if m.size == 0 {
		return chans
	}
	chans = make([]chan tuple[K, mapcommon.MapItem[T]], m.size)
	wg := sync.WaitGroup{}
	wg.Add(m.size)
	for index, shard := range m.shards {
		go func(index int, shard *mapshared.Shared[K, T]) {
			shard.RLock()
			chans[index] = make(chan tuple[K, mapcommon.MapItem[T]], shard.CountItems())
			wg.Done()
			for key, val := range shard.GetItems() {
				chans[index] <- tuple[K, mapcommon.MapItem[T]]{key, val}
			}
			shard.RUnlock()
			close(chans[index])
		}(index, shard)
	}
	wg.Wait()
	return chans
}

// 多线程组装
func fanIn[K comparable, T any](chans []chan tuple[K, T], out chan tuple[K, T]) {
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan tuple[K, T]) {
			for t := range ch {
				out <- t
			}
			wg.Done()
		}(ch)
	}
	wg.Wait()
	close(out)
}
