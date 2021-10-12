package cache

import "sync"

type GodisCache struct {
	coreMap map[string]Frame
	mutex   sync.RWMutex
	Current Overview
}

func (g *GodisCache) Set(key string, value *Frame) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

}
