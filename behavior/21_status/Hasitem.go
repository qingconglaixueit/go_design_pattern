// @Author Bing 
// @Desc
package main

import (
	"errors"
	"fmt"
)

type HasItem struct {
	Machine *ShopMachine
}

func (h *HasItem) AddItem(num int) error {
	h.Machine.IncItem(num)
	return nil
}
func (h *HasItem) ReqItem() error {
	if h.Machine.ItemCnt == 0 {
		h.Machine.SetState(h.Machine.NoItemState)
		return errors.New(("no item ... "))
	}
	fmt.Println("req item ... ")
	h.Machine.SetState(h.Machine.ReqItemState)
	return nil
}
func (h *HasItem) InsertMoney(int) error {
	return errors.New("please select item ... ")
}
func (h *HasItem) GetItem() error {
	return errors.New("please select item ... ")
}
