// @Author Bing 
// @Date 2023/3/27 16:08:00 
// @Desc
package main

import "fmt"

func main() {

	nb,_ := GetShopFac("NB")

	nike,_ := GetShopFac("Nike")

	nbCloth := nb.MakeCloth()
	fmt.Printf("cloth brand:[%s],color:[%s],size[%s]\n\n",nbCloth.GetBrand(),nbCloth.GetColor(),nbCloth.GetSize())

	nikeShor := nike.MakeShoe()
	fmt.Printf("shoe brand:[%s],color:[%s],size[%d]\n",nikeShor.GetBrand(),nikeShor.GetColor(),nikeShor.GetSize())
}
