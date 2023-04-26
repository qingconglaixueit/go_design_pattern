// @Author Bing 
// @Desc
package main

import "fmt"

type Contract struct {
	next IDepartment
}

func (f *Contract) Execute(e *Employee) {
	if e.ContractDone{
		fmt.Errorf("Contract has serverd ... ")
	}
	fmt.Println("start Contract ... ")
	e.ContractDone = true
	f.next.Execute(e)
}
func (f *Contract) SetNext(d IDepartment) {
	f.next = d
}