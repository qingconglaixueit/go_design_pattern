// @Author Bing 
// @Desc
package main

import "fmt"

type File struct {
	Name string
}

func (f *File) Search(content string) {
	fmt.Printf("search file name :[%s] ,content:[%s]\n", f.Name, content)
}

func (f *File) GetName() string {
	return f.Name
}
