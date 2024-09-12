package mem_cache

import (
	"context"
	"github.com/puresnr/pgo/algo"
	"github.com/puresnr/pgo/go_safe"
	"sync"
	"time"
)

type memCache[K comparable, V any] struct {
	locker       *sync.RWMutex
	cache        map[K]V
	funcNewCache func(context.Context, K) (V, error)
}

func (m *memCache[K, V]) Get(ctx context.Context, key K) (V, error) {
	m.locker.RLock()
	v, ok := m.cache[key]
	m.locker.RUnlock()
	if ok {
		return v, nil
	}

	var err error
	if v, err = m.funcNewCache(ctx, key); err != nil {
		return v, err
	}

	m.locker.Lock()
	m.cache[key] = v
	m.locker.Unlock()

	return v, nil
}

func (m *memCache[K, V]) del() {
	m.locker.Lock()
	c, dc := 0, len(m.cache)/3
	if c != dc {
		for k, _ := range m.cache {
			delete(m.cache, k)
			c++
			if c == dc {
				break
			}
		}
	}
	m.locker.Unlock()
}

func New[K comparable, V any](funcNewCache func(context.Context, K) (V, error), delDura ...uint) *memCache[K, V] {
	mc := &memCache[K, V]{locker: new(sync.RWMutex), cache: make(map[K]V), funcNewCache: funcNewCache}

	if !algo.Empty_slice(delDura) {
		go_safe.GoR(func() {
			ticker := time.NewTicker(time.Duration(delDura[0]) * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					mc.del()
				}
			}
		})
	}

	return mc
}
