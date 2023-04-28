// @Author Bing 
// @Desc
package main

type Pearl struct {
	MilkTea Imilktea
}

func (n *Pearl) GetPrice() int {
	return n.MilkTea.GetPrice() + 3
}
