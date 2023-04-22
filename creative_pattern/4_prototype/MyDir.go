// @Author Bing 
// @Desc
package main

import "fmt"

type MyDir struct {
	Name  string
	Child []INode
}

func (d *MyDir) Show(sep string) {
	fmt.Println(sep + d.Name)
	for _, ch := range d.Child {
		ch.Show(sep + sep)
	}
}

func (d *MyDir) Clone() INode {
	cloneChild := make([]INode, len(d.Child))

	for index, ch := range d.Child {
		cloneChild[index] = ch.Clone()
	}

	return &MyDir{Name: d.Name + "_clone", Child: cloneChild}
}
