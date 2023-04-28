// @Author Bing 
// @Desc
package main


func main(){

	c := &Client{}

	m := &Mac{}
	c.InsertUsbToComputer(m)

	// type C 的接口 通过 适配器 转换成可以通过 USB 插入到 电脑中
	oth := &Other{}
	ad := &Adapter{OtherSvr: oth}
	c.InsertUsbToComputer(ad)

}

