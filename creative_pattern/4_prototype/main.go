// @Author Bing 
// @Desc
package main
// @Title main
// @Description  能够满足递归的结构，目录和文件，都是同样的结构
func main() {
	f1 := &MyFile{"文件1"}
	f2 := &MyFile{"文件2"}

	dir1 := &MyDir{Name: "目录1", Child: []INode{f1}}
	dir2 := &MyDir{Name: "目录2", Child: []INode{f2, dir1, f1}}

	dir2.Show("  ")
}
