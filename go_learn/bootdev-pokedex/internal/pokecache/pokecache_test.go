package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	// Verify the key exists before reaping
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	// Verify the key is removed after reaping
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}

func TestCacheConcurrency(t *testing.T) {
	const interval = 1 * time.Second
	cache := NewCache(interval)

	key := "https://example.com"
	val := []byte("concurrency-test")
	cache.Add(key, val)

	done := make(chan bool)

	// Start multiple goroutines to access the cache concurrently
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				_, ok := cache.Get(key)
				if !ok {
					t.Errorf("expected to find key during concurrency test")
				}
			}
			done <- true
		}()
	}

	// Wait for all goroutines to finish
	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestCacheExpiration(t *testing.T) {
	const interval = 100 * time.Millisecond
	cache := NewCache(interval)
	key := "https://example.com"
	val := []byte("expiry-test")
	cache.Add(key, val)

	time.Sleep(interval / 2)

	// Verify the key still exists before expiration
	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key before expiration")
		return
	}

	time.Sleep(interval)

	// Verify the key is removed after expiration
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected key to be expired")
	}
}
