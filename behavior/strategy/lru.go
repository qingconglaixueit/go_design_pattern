// @Author Bing 
// @Desc
package main

import "fmt"

type Lru struct {
}

func (f *Lru) Evict(c *Cache) {
	fmt.Println("Lru evict cache")
}
