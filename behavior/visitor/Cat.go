// @Author Bing 
// @Desc
package main

import "fmt"

type Cat struct{}

func (c *Cat) GetType() {
	fmt.Println(" cat ... ")
}
func (c *Cat) Accept(v IVisitor) {
	v.VsCat(c)
}
