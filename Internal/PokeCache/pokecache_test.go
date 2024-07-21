package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestGetCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	} {
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
	
		if !ok {
			t.Errorf("%v not found", cas.inputKey)
			continue
		}
		
		if string(actual) != string(cas.inputVal) {
			t.Error("values doesn't match")
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key1 := "key1"
	cache.Add(key1, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key1)

	if ok {
		t.Error("key '%s' should have been reaped", key1)
	}


}

