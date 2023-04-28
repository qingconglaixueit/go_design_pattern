// @Author Bing 
// @Desc
package main

import "fmt"

type AreaVs struct {
}

func (a *AreaVs) VsDog(D *Dog) {
	fmt.Println(" cal dog area ... ")
}
func (a *AreaVs) VsCat(C *Cat) {
	fmt.Println(" cal cat area ... ")
}
func (a *AreaVs) VsMonkey(M *Monkey) {
	fmt.Println(" cal monkey area ... ")
}
