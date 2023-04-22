// @Author Bing 
// @Desc
package main

import "fmt"

type Lfu struct{

}

func (f *Lfu)Evict(c * Cache){
	fmt.Println("Lfu evict cache")
}
