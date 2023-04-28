// @Author Bing 
// @Desc
package main

type OnCmd struct {
	Dev Device
}

func (o *OnCmd) Execute() {
	o.Dev.on()
}
