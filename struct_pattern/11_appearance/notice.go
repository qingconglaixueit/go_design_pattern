// @Author Bing 
// @Desc
package main

import "fmt"

type NoticeSys struct {
	Prefix string
}

func NewNotice(prefix string) *NoticeSys {
	return &NoticeSys{Prefix: prefix}
}

func (n *NoticeSys) SendNotice(msg string) {
	fmt.Printf("[%s]:%s \n", n.Prefix, msg)
}
