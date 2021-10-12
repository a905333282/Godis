package cache

import "sync"

type GodisCache struct {
	coreMap map[string]Frame
	mutex sync.RWMutex
	Current Overview
}


Set(string, *Frame) error
Get(string) (*Frame, error)
Del(string) error
Overview() Overview


func (g *GodisCache) Set(key string, value *Frame)  {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	
}