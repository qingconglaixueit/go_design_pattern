// @Author Bing 
// @Desc
package main

import "fmt"

type Teacher struct {
	Mgr IOrderMgr
}

func (t *Teacher) Enter() {
	if !t.Mgr.CanEnter(t){
		fmt.Println(" [Block]queue block , teacher enter ... ")
		return
	}
	fmt.Println(" teacher enter ... ")
}
func (t *Teacher) Leave() {
	fmt.Println(" teacher leave ... ")
	t.Mgr.NoticePeople()
}
func (t *Teacher) PermitEnter() {
	fmt.Println(" PermitEnter  ... ")
	t.Enter()
}
