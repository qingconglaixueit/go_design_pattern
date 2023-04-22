// @Author Bing 
// @Desc
package main
// IBuild  建一座建筑，大体会经历如下步骤
// 1. 搭建框架
// 2. 装修
// 3. 安装门
// 4. 安装床
// ...
type IBuild interface{
	SetFrame()
	SetStyle()
	SetDoor()
	SetBed()
	BuildHouse()*House
}

func GetBuild(bType string)IBuild{
	if bType == "ice"{
		return NewIce()
	}
	if bType == "wood"{
		return NewWood()
	}
	return nil
}