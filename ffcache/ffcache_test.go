package ffCache

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestGetter(t *testing.T) {
	var f Getter = GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(expect, v) {
		t.Errorf("callback failed")
	}
}

var db = map[string]string{
	"Tom":  "630",
	"Jack": "720",
	"Sam":  "560",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))

	ff := NewGroup("score", 2<<10, GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			if _, ok := loadCounts[key]; ok {
				loadCounts[key] = 0
			}
			loadCounts[key]++
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	for k, v := range db {
		if view, err := ff.Get(k); err != nil || view.String() != v {
			t.Fatalf("failed to get value of Tom")
		}
		if _, err := ff.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("cache %s misss", k)
		}
	}

	if view, err := ff.Get("unknown"); err == nil {
		t.Fatalf("the value of unknown be empty, but %s get", view)
	}
}
