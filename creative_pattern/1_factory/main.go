// @Author Bing
// @Desc
package main

import "fmt"

func main() {
	bmw, _ := getCar("Bmw")

	benz ,_ := getCar("Benz")

	fmt.Println(bmw.GetName(),bmw.GetColor())
	bmw.Run()


	fmt.Println(benz.GetName(),benz.GetColor())
	benz.Run()
}
