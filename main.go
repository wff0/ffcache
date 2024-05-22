package main

import (
	ffCache "ffCache/ffcache"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "452",
	"Jack": "444",
	"Sam":  "555",
}

func main() {
	ffCache.NewGroup("scores", 2<<10, ffCache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:8888"
	peers := ffCache.NewHTTPPool(addr)
	log.Println("ffcache is running at", addr)
	log.Fatal(http.ListenAndServe(addr, peers))
}
