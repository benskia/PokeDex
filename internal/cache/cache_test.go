package cache

import "testing"

func TestCache(t *testing.T) {
	cache := NewCache(5)
	cases := []struct {
		inputKey string
		inputVal []byte
		expected bool
	}{
		{
			inputKey: "poketest",
			expected: true,
		},
		{
			inputKey: "pokewhat",
			expected: true,
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, nil)
	}

	for _, cs := range cases {
		_, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("Key %v not found in cache.", cs.inputKey)
		}
	}
}

func TestCacheInvalid(t *testing.T) {
	cache := NewCache(5)
	cases := []struct {
		inputKey string
		inputVal []byte
		expected bool
	}{
		{
			inputKey: "poketest",
			expected: false,
		},
		{
			inputKey: "pokewhat",
			expected: false,
		},
	}

	for _, cs := range cases {
		_, ok := cache.Get(cs.inputKey)
		if ok {
			t.Errorf("Non-existent key %v was found in cache.", cs.inputKey)
		}
	}
}
