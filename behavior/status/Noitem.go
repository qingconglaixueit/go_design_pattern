// @Author Bing 
// @Desc
package main

import "errors"

type NoItem struct {
	Machine *ShopMachine
}

func (n *NoItem) AddItem(num int) error {
	n.Machine.IncItem(num)
	n.Machine.SetState(n.Machine.HasItemState)

	return nil
}
func (n *NoItem) ReqItem() error {
	return errors.New("no item ... ")
}
func (n *NoItem) InsertMoney(int) error {
	return errors.New("no item ... ")
}
func (n *NoItem) GetItem() error {
	return errors.New("no item ... ")
}
