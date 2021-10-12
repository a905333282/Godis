package cache

type Cache interface {
	Set(string, *Frame) error
	Get(string) (*Frame, error)
	Del(string) error
	Overview() Overview
}
