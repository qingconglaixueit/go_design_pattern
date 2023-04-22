// @Author Bing 
// @Desc
package main

import "fmt"
// @Title main
// @Description  迭代器模式，隐藏实际的数据结构
func main() {
	ua := UserArray{users: []*User{&User{Name: "xiaozhu1", Age: 10}, &User{Name: "xiaozhu2", Age: 18}}}

	it := ua.createIterator()

	for it.HasNext() {
		fmt.Println(it.GetNext())
	}
}
