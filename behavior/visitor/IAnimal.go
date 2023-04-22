// @Author Bing 
// @Desc
package main

type IAnimal interface {
	GetType()
	Accept(v IVisitor)
}
