package cache

import "sync"

type GodisCache struct {
	coreMap map[string]*Frame
	mutex   sync.RWMutex
	Current *Overview
}

func (g *GodisCache) Set(key string, value *Frame) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	g.coreMap[key] = value
	g.Current.Add(value)

	return nil
}

func (g *GodisCache) Get(key string) (*Frame, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	return g.coreMap[key], nil
}

func (g *GodisCache) Del(key string) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	value, exist := g.coreMap[key]
	if exist {
		delete(g.coreMap, key)
		g.Current.Del(value)
	}
	return nil
}

func (g *GodisCache) Overview() *Overview {
	return g.Current
}

func NewGodisCache() *GodisCache {
	return &GodisCache{make(map[string]*Frame), sync.RWMutex{}, NewOverview()}
}
