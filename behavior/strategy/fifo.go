// @Author Bing 
// @Desc
package main

import "fmt"

type Fifo struct{

}

func (f *Fifo)Evict(c * Cache){
	fmt.Println("fifo evict cache")
}