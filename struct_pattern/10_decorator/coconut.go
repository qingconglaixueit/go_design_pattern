// @Author Bing 
// @Desc
package main

type Coconut struct {
	MilkTea Imilktea
}

func (n *Coconut) GetPrice() int {
	return n.MilkTea.GetPrice() + 5
}
