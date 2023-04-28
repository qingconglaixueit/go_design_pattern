// @Author Bing 
// @Desc
package main

import "fmt"

type Front struct {
	next IDepartment
}

func (f *Front) Execute(e *Employee) {
	if e.FrontDone{
		fmt.Errorf("Front has serverd ... ")
	}
	fmt.Println("start front ... ")
	e.FrontDone = true
	f.next.Execute(e)

}
func (f *Front) SetNext(d IDepartment) {
	f.next = d
}
