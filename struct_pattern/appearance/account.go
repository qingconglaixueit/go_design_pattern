// @Author Bing 
// @Desc
package main

type Account struct {
	Name string
}

func NewAccount(name string) *Account {
	return &Account{Name: name}
}

func (a *Account) CheckAccount(name string) bool {
	return name == a.Name
}
