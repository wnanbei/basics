package cache

import (
	"sync"

	"github.com/maypok86/otter"
)

const (
	// globalLimit 全局缓存池大小
	globalLimit int = 100_000
)

var globalOnce sync.Once

// global 全局缓存
var global otter.CacheWithVariableTTL[string, any]

// Global 获取全局缓存池
func Global() otter.CacheWithVariableTTL[string, any] {
	globalOnce.Do(func() {
		global = New[string, any](globalLimit)
	})

	return global
}

// New 创建缓存池
func New[K comparable, V any](limit int) otter.CacheWithVariableTTL[K, V] {
	cache, _ := otter.MustBuilder[K, V](limit).WithVariableTTL().Build()
	return cache
}
