// @Author Bing 
// @Date 2023/3/27 16:08:00 
// @Desc
package main

import "fmt"
// 一个大卖场，有衣服，有 鞋子，有 NB 品牌，有 Nike 品牌，此时我们就可以使用抽象工厂，抽象品牌，品牌下面有多种产品
func main() {

	nb,_ := GetShopFac("NB")

	nike,_ := GetShopFac("Nike")

	nbCloth := nb.MakeCloth()
	fmt.Printf("cloth brand:[%s],color:[%s],size[%s]\n\n",nbCloth.GetBrand(),nbCloth.GetColor(),nbCloth.GetSize())

	nikeShor := nike.MakeShoe()
	fmt.Printf("shoe brand:[%s],color:[%s],size[%d]\n",nikeShor.GetBrand(),nikeShor.GetColor(),nikeShor.GetSize())
}
