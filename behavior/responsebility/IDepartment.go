// @Author Bing 
// @Desc
package main

type IDepartment interface {
	Execute(e *Employee)
	SetNext(d IDepartment)
}
