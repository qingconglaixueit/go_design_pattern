// @Author Bing 
// @Desc
package main

import "fmt"

type MyFile struct {
	Name string
}

func(f *MyFile)Show(sep string){
	fmt.Println(sep + f.Name)
}

func(f *MyFile)Clone()INode{
	return &MyFile{Name: f.Name + "_clone"}
}