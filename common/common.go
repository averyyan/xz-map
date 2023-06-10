package mapcommon

import (
	"fmt"
	"time"
)

type MapStringer interface {
	comparable
	fmt.Stringer
}

// 缓存数据接口
type MapItem[T any] interface {
	GetValue() T                             // 获取缓存值
	VerifyExpiration(unixNano int64) bool    // 判断是否过期
	VerifyTimeDuration(d time.Duration) bool // 判断是否超过时间
	DeleteHandler() error                    // 删除执行函数
}
