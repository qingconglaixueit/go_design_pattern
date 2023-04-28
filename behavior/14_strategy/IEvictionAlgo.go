// @Author Bing 
// @Desc
package main

type EvictionAlgo interface{
	Evict(c * Cache)
}
