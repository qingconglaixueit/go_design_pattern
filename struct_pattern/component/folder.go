// @Author Bing
// @Desc
package main

import "fmt"

type Folder struct {
	Com  []Icomponent
	Name string
}

func (f *Folder) Search(content string) {
	fmt.Printf("strat search folder , name:[%s] content:[%s] \n", f.Name, content)
	for _, v := range f.Com {
		v.Search(content)
	}

}
func (f *Folder) AddComponent(c Icomponent) {
	f.Com = append(f.Com, c)
}
