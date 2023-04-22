// @Author Bing 
// @Desc
package main

import (
	"errors"
	"fmt"
)

type ReqItem struct {
	Machine *ShopMachine
}

func (r *ReqItem) AddItem(int) error {
	return errors.New("item is processing ... ")
}
func (r *ReqItem) ReqItem() error {
	return errors.New("in reqitem ...")
}
func (r *ReqItem) InsertMoney(money int) error {
	fmt.Println("insert money ", money)
	if money < r.Machine.ItemPrice {
		return errors.New("not enough money ... ")
	}
	r.Machine.SetState(r.Machine.HasMoney)
	return nil
}
func (r *ReqItem) GetItem() error {
	return errors.New("please insert money ... ")
}
