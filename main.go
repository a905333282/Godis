package main

import (
	"Godis/cache"
	"fmt"
)

func main() {
	g := cache.NewGodisCache()
	//g.Set("1", cache.NewFrame{"1",[]byte("")})
	fmt.Println(g.Overview())
}
