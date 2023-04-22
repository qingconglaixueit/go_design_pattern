// @Author Bing 
// @Desc
package main

type Item interface{
	Register(Observe)
	Deregister(Observe)
	NoticeAll()
}