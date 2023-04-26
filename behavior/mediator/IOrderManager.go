// @Author Bing 
// @Desc
package main

type IOrderMgr interface{
	NoticePeople()
	CanEnter(people IPeople)bool
}
