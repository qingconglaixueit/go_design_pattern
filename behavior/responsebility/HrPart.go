// @Author Bing 
// @Desc
package main

import "fmt"

type HrPart struct {
	next IDepartment
}

func (f *HrPart) Execute(e *Employee) {
	if e.HrDone{
		fmt.Errorf("HrPart has serverd ... ")
	}
	fmt.Println("start HrPart ... ")
	e.HrDone = true
	f.next.Execute(e)
}
func (f *HrPart) SetNext(d IDepartment) {
	f.next = d
}