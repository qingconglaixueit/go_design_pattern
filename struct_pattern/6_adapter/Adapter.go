// @Author Bing 
// @Desc
package main

import "fmt"

type Adapter struct{
	OtherSvr *Other
}

func (a *Adapter)InsertUsb(){
	fmt.Println("**********  Convert other interface to Usb ... *********")
	a.OtherSvr.InsertTypeC()
}