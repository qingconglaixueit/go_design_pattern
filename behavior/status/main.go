// @Author Bing 
// @Desc
package main

import "fmt"

// @Title main
// @Description  售卖机暂时只售卖一种从产品，有 4 种状态，没货，有货，正在请求商品，已投币
func main() {

	shop := newSHopMachine(2, 100)

	shop.AddItem(1)

	shop.ReqItem()

	fmt.Println(shop.InsertMoney(10))

	fmt.Println(shop.GetItem())

	shop.InsertMoney(100)

	shop.GetItem()

	fmt.Println("----------------------------------")

	shop.ReqItem()

	fmt.Println(shop.GetItem())

	shop.InsertMoney(200)
	shop.GetItem()

}
