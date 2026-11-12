package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("a", []byte("hello"))
	val, ok := cache.Get("a")
	if !ok {
		t.Errorf("key not found")
	}
	if string(val) != "hello" {
		t.Errorf("expected hello, not %s", string(val))
	}
}

func TestGetMissing(t *testing.T) {
	cache := NewCache(5 * time.Second)
	val, ok := cache.Get("nothing")
	if ok{
		t.Errorf("key mustn't to be found")
	}
	if val != nil{
		t.Errorf("expected nil")
	}
}

func TestReplace(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("a", []byte("old"))
	cache.Add("b", []byte("other"))
	cache.Add("a", []byte("new"))

	val, _ := cache.Get("a")
	if string(val) != "new"{
		t.Errorf("expected new, not %s", string(val))
	}

	val, _ = cache.Get("b")
	if string(val) != "other" {
		t.Errorf("expected other, not %s", string(val))
	}
}

func TestReap(t *testing.T){
	interval := 50 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("a", []byte("value"))
	time.Sleep(10*time.Millisecond)
	_, ok := cache.Get("a")
	if !ok {
		t.Errorf("expected entry to still exist")
	}

	time.Sleep(60*time.Millisecond)
	_, ok = cache.Get("a")
	if ok {
		t.Errorf("expected entry to be reaped")
	}
}