// @Author Bing 
// @Desc
package main

type Interator interface {
	HasNext() bool
	GetNext() *User
}
