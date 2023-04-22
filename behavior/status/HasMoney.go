// @Author Bing 
// @Desc
package main

import (
	"errors"
	"fmt"
)

type HasMoney struct {
	Machine *ShopMachine
}

func (h *HasMoney) AddItem(int) error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) ReqItem() error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) InsertMoney(int) error {
	return errors.New("item is processing ... ")
}
func (h *HasMoney) GetItem() error {
	fmt.Println("get item ...")
	h.Machine.ItemCnt -= 1
	if h.Machine.ItemCnt == 0 {
		h.Machine.SetState(h.Machine.NoItemState)
	} else {
		h.Machine.SetState(h.Machine.HasItemState)
	}
	return nil
}
