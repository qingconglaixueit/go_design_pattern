// @Author Bing 
// @Desc
package main

type IState interface {
	AddItem(int) error
	ReqItem() error
	InsertMoney(int) error
	GetItem() error
}
