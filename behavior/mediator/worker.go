// @Author Bing 
// @Desc
package main

import "fmt"

type Worker struct {
	Mgr IOrderMgr
}

func (t *Worker) Enter() {
	if !t.Mgr.CanEnter(t){
		fmt.Println(" [Block]queue block , Worker enter ... ")
		return
	}
	fmt.Println(" Worker enter ... ")
}
func (t *Worker) Leave() {
	fmt.Println(" Worker leave ... ")
	t.Mgr.NoticePeople()
}
func (t *Worker) PermitEnter() {
	fmt.Println(" PermitEnter  ... ")
	t.Enter()
}
