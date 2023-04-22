// @Author Bing 
// @Desc
package main

import "fmt"

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
