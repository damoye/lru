package lru

import (
	"testing"
)

func TestCache(t *testing.T) {
	cache := newCache(2)

	// {}
	value, ok := cache.Get("k1")
	if ok {
		t.Fatal("got true want false")
	}

	// {"k1": "v1"}
	cache.Add("k1", "v1")

	value, ok = cache.Get("k1")
	if !ok {
		t.Fatal("got false want true")
	}
	if value != "v1" {
		t.Fatalf("got %q want %q", value, "v1")
	}

	value, ok = cache.Get("k2")
	if ok {
		t.Fatal("got true want false")
	}

	// {"k2": "v2", "k1": "v1"}
	cache.Add("k2", "v2")
	value, ok = cache.Get("k2")
	if !ok {
		t.Fatal("got false want true")
	}
	if value != "v2" {
		t.Fatalf("got %q want %q", value, "v2")
	}

	// {"k3": "v3", "k2": "v2"}
	cache.Add("k3", "v3")
	value, ok = cache.Get("k3")
	if !ok {
		t.Fatal("got false want true")
	}
	if value != "v3" {
		t.Fatalf("got %q want %q", value, "v3")
	}

	value, ok = cache.Get("k1")
	if ok {
		t.Fatal("got true want false")
	}
}
