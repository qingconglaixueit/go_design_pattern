// @Author Bing 
// @Desc
package main

import "fmt"

type ColorVs struct {
}

func (a *ColorVs) VsDog(D *Dog) {
	fmt.Println(" get dog color ... ")
}
func (a *ColorVs) VsCat(C *Cat) {
	fmt.Println(" get cat color ... ")
}
func (a *ColorVs) VsMonkey(M *Monkey) {
	fmt.Println(" get monkey color ... ")
}
