// @Author Bing 
// @Desc
package main

import "fmt"

// 喝奶茶，有原味的，有加珍珠的，有加椰果的，有加奶的 ，各有各的价格
// 那么 奶茶 和 珍珠，椰果，奶 是相互独立的，各自是一个独立的对象
// 我们可以在珍珠对象中，放入奶茶对象，进而达到 奶茶中加珍珠，加椰果，加奶  达到不同的价格
func main(){

	// 一杯原味 naixue 奶茶
	mt := &Naixue{}
	fmt.Println("原味 ： ",  mt.GetPrice())


	// 加了珍珠的 naixue
	p := &Pearl{MilkTea: mt}
	fmt.Println("原味+珍珠 ： ",p.GetPrice())

	// 加了椰果的 naixue
	c := &Coconut{MilkTea: mt}
	fmt.Println("原味+椰果 ： ",c.GetPrice())


	// 加了牛奶的 naixue

	m := &Milk{MilkTea: mt}
	fmt.Println("原味+牛奶 ： ",m.GetPrice())



}
