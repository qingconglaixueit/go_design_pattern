// @Author Bing 
// @Desc
package main

type Cache struct {
	Storage map[string]string
	E       EvictionAlgo
	Cap     int
	MaxCap  int
}

func initCache(e EvictionAlgo)*Cache{
	return &Cache{
		Storage: make(map[string]string,0),
		E:  e,
		Cap: 0,
		MaxCap: 3,
	}
}

func (c *Cache)Add(key string,value string){
	if c.Cap == c.MaxCap {
		c.Evict()
	}
	c.Cap++
	c.Storage[key] = value
}
func (c * Cache)SetEvi(e EvictionAlgo){
	c.E = e
}

func (c * Cache)Evict(){
	c.E.Evict(c)
	c.Cap--
}