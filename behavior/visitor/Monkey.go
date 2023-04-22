// @Author Bing 
// @Desc
package main

import "fmt"

type Monkey struct{}


func (c *Monkey) GetType() {
	fmt.Println(" Monkey ... ")
}
func (c *Monkey) Accept(v IVisitor) {
	v.VsMonkey(c)
}