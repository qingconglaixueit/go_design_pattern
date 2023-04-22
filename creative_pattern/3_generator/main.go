// @Author Bing 
// @Desc
package main

import "fmt"
//  建一座建筑，无论你是冰室，还是住房，大体会经历如下步骤
// 1. 搭建框架
// 2. 装修
// 3. 安装门
// 4. 安装床
func main() {
	b := GetBuild("ice")
	if b == nil {
		fmt.Println("not support bType ...")
		return
	}
	m := NewManager(b)
	iceHouse := m.GetHouse()
	fmt.Println("iceHouse == ", iceHouse)


	b2 := GetBuild("wood")
	if b2 == nil {
		fmt.Println("not support bType ...")
		return
	}
	m2 := NewManager(b2)
	woodHouse := m2.GetHouse()
	fmt.Println("woodHouse == ", woodHouse)


	m2.SetWorker(b)
	iihouse := m2.GetHouse()
	fmt.Println("iihouse == ", iihouse)

}
