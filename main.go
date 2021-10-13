package main

import (
	"Godis/cache"
	"fmt"
)

func main() {
	g := cache.NewGodisCache()
	g.Set("1", cache.Frame{"1"})
	fmt.Println(g.Overview())
}
