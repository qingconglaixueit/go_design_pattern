// @Author Bing 
// @Desc
package main

import "fmt"

type VolumeVs struct {
}

func (a *VolumeVs) VsDog(D *Dog) {
	fmt.Println(" cal dog volume ... ")
}
func (a *VolumeVs) VsCat(C *Cat) {
	fmt.Println(" cal cat volume ... ")
}
func (a *VolumeVs) VsMonkey(M *Monkey) {
	fmt.Println(" cal monkey volume ... ")
}
