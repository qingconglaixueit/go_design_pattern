// @Author Bing 
// @Desc
package main

type Milk struct {
	MilkTea Imilktea
}

func (n *Milk) GetPrice() int {
	return n.MilkTea.GetPrice() + 10
}
