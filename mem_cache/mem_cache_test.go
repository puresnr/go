package mem_cache

import (
	"context"
	"testing"
	"time"
)

func TestMemCache(t *testing.T) {
	cache := New[int, int64](func(ctx context.Context, i int) (int64, error) {
		return time.Now().Unix(), nil
	}, 5)

	time.Sleep(6 * time.Second)

	tnow := time.Now().Unix()
	t1, _ := cache.Get(context.Background(), 1)
	if t1 != tnow {
		t.Fatal("cache not working")
	}

	time.Sleep(6 * time.Second)

	if tt1, _ := cache.Get(context.Background(), 1); tt1 != t1 {
		t.Fatal("cache not working")
	}

	t2, _ := cache.Get(context.Background(), 2)
	time.Sleep(time.Second)
	t3, _ := cache.Get(context.Background(), 3)

	time.Sleep(5 * time.Second)

	tt1, _ := cache.Get(context.Background(), 1)
	tt2, _ := cache.Get(context.Background(), 2)
	tt3, _ := cache.Get(context.Background(), 3)

	if !((tt1 == t1 && tt2 == t2 && tt3 != t3) ||
		(tt1 == t1 && tt2 != t2 && tt3 == t3) ||
		(tt1 != t1 && tt2 == t2 && tt3 == t3)) {
		t.Fatal("cache not working")
	}
}
