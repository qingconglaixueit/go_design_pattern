// @Author Bing 
// @Desc
package main

type OffCmd struct{
	Dev Device
}

func (o *OffCmd)Execute(){
	o.Dev.off()
}