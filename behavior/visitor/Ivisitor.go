// @Author Bing 
// @Desc
package main

type IVisitor interface {
	VsDog(D *Dog)
	VsCat(C *Cat)
	VsMonkey(M *Monkey)
}
