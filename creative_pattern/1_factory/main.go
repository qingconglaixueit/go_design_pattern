// @Author Bing
// @Desc
package main

import "fmt"
// 写一个工厂，可以生产车，目前这个工厂可以生产 bmw 和 benz ，那么我们需要获取一辆车的时候，就直接去车厂拿车即可
func main() {
	bmw, _ := getCar("Bmw")

	benz ,_ := getCar("Benz")

	fmt.Println(bmw.GetName(),bmw.GetColor())
	bmw.Run()


	fmt.Println(benz.GetName(),benz.GetColor())
	benz.Run()
}
