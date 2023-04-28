// @Author Bing 
// @Desc
package main

import "fmt"

// 可以使用它将对象组合成树状结构， 并且能像使用独立对象一样使用它们
// 例如咱们的目录树中有 目录，有 文件， 目录可以进行查询，文件也是可以进行查询，将他们组合起来，他们各自仍然是独立的对象
//folder-2
//	file-1
//	file-2
//	file-3
//	folder-1
//		file-1
//		file-2
//		file-3

func main() {
	f1 := &File{Name: "file-1"}
	f2 := &File{Name: "file-2"}
	f3 := &File{Name: "file-3"}

	fo1 := &Folder{Name: "folder-1"}
	fo2 := &Folder{Name: "folder-2"}

	fo1.AddComponent(f1)
	fo1.AddComponent(f2)
	fo1.AddComponent(f3)


	fo2.AddComponent(f1)
	fo2.AddComponent(f2)
	fo2.AddComponent(f3)
	fo2.AddComponent(fo1)

	fo2.Search("xiaozhu")

	fmt.Printf("-------------------------------\n")

	fo1.Search("xiaopang")


}
