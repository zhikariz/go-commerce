package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	lru "github.com/hnlq715/golang-lru"
)

func InitLRUCache(size int) (*lru.ARCCache, error) {
	arcCache, err := lru.NewARC(size)
	if err != nil {
		return nil, err
	}
	return arcCache, err
}

type LRUCacheable interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, result interface{}) error
}

type arcCacheable struct {
	arc *lru.ARCCache
}

func NewARCCacheable(arc *lru.ARCCache) *arcCacheable {
	return &arcCacheable{
		arc: arc,
	}
}

func (c *arcCacheable) Set(key string, value interface{}, expiration time.Duration) error {
	marshalledJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}
	c.arc.AddEx(key, string(marshalledJSON), expiration)
	return nil
}

func (c *arcCacheable) Get(key string, result interface{}) error {
	val, ok := c.arc.Get(key)
	if !ok {
		return errors.New("key not found / expired")
	}
	fmt.Println("val", val)
	fmt.Println("ok", ok)
	value := fmt.Sprintf("%v", val)
	if val != "" {
		err := json.Unmarshal([]byte(value), &result)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
