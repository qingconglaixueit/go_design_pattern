// @Author Bing 
// @Desc
package main

import "fmt"

type OwnPart struct {
	next IDepartment
}

func (f *OwnPart) Execute(e * Employee) {
	if e.EnterSpecPartDone{
		fmt.Errorf("OwnPart has serverd ... ")
	}
	fmt.Println("start OwnPart ... ")
	e.EnterSpecPartDone = true
}
func (f *OwnPart) SetNext(d IDepartment) {
	fmt.Printf("employee has been entered department ...")
}
