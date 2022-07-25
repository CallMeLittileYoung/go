package chapter5

import "sync"

var (
	mu      sync.Mutex
	mapping = make(map[string]int)
)

func LookUp(key string) int {
	mu.Lock()
	defer mu.Unlock()
	i := mapping[key]
	return i
}

type Cache struct {
	sync.Mutex
	mapping map[string]int
}

func (c Cache) GetCache(key string) int {
	c.Lock()
	defer c.Unlock()
	return mapping[key]
}
