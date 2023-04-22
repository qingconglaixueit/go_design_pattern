// @Author Bing 
// @Desc
package main

import "fmt"

type Dog struct{}

func (c *Dog) GetType() {
	fmt.Println(" Dog ... ")
}
func (c *Dog) Accept(v IVisitor) {
	v.VsDog(c)
}