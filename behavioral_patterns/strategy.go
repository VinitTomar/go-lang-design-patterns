package behavioral_patterns

import "fmt"

type evictionAlgorithm interface {
	evict(c *cache)
}

type fifoEvict struct {}

func (e *fifoEvict) evict(c *cache) {
	fmt.Println("Evicting by FIFO.")
}

type lruEvict struct {}

func (e *lruEvict) evict(c *cache) {
	fmt.Println("Evicting by LRU.")
}

type lfuEvict struct {}

func (e *lfuEvict) evict(c *cache) {
	fmt.Println("Eviction by LFU.")
}

type cache struct {
	storage map[string]string
	evictionAlgo evictionAlgorithm
	capacity int
	maxCapacity int
}

func initCache(e evictionAlgorithm) *cache {
	return &cache{
		evictionAlgo: e,
		storage: make(map[string]string),
		capacity: 0,
		maxCapacity: 2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgorithm) {
	c.evictionAlgo = e
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}

	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) string {
	value := c.storage[key]
	delete(c.storage, key)
	return value
}

func StrategyPattern() {
	lfu := &lfuEvict{}
	cache := initCache(lfu)

	cache.add("1", "one")
	cache.add("2", "two")

	cache.add("3", "three")

	lru := &lruEvict{}

	cache.setEvictionAlgo(lru)

	cache.add("4", "four")
	
	fifo := &fifoEvict{}

	cache.setEvictionAlgo(fifo)

	cache.add("5", "five")

	fmt.Printf("Value for key %v is %v\n", "5", cache.get("5"))
}