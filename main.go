package main

import "fmt"
import "Godis/cache"

func main() {
	n := cache.NewFrame("1", []byte("str"))
	fmt.Println(n.GetValueSize())
}
