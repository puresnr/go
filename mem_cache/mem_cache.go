package mem_cache

import (
	"github.com/puresnr/pgo/algo"
	"github.com/puresnr/pgo/gosafe"
	"sync"
	"time"
)

type memCache[K comparable, V, P any] struct {
	locker  *sync.RWMutex
	cache   map[K]V
	funcGen func(K, P) (V, error)
}

func (m *memCache[K, V, P]) Get(key K, param P) (V, error) {
	m.locker.RLock()
	v, ok := m.cache[key]
	m.locker.RUnlock()
	if ok {
		return v, nil
	}

	var err error
	if v, err = m.funcGen(key, param); err != nil {
		return v, err
	}

	m.locker.Lock()
	if v1, ok := m.cache[key]; ok {
		v = v1
	} else {
		m.cache[key] = v
	}
	m.locker.Unlock()

	return v, nil
}

func (m *memCache[K, V, P]) del() {
	m.locker.Lock()
	c, dc := 0, len(m.cache)/3
	if c != dc {
		for k := range m.cache {
			delete(m.cache, k)
			c++
			if c == dc {
				break
			}
		}
	}
	m.locker.Unlock()
}

func New[K comparable, V, P any](funcGen func(K, P) (V, error), delDura ...uint) *memCache[K, V, P] {
	mc := &memCache[K, V, P]{locker: new(sync.RWMutex), cache: make(map[K]V), funcGen: funcGen}

	if !algo.Empty_slice(delDura) {
		gosafe.GoR(func() {
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
